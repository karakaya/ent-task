package pkg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionState string

const (
	StateWin  TransactionState = "win"
	StateLose TransactionState = "lose"
)

type UserTransaction struct {
	UserId        uint64           `json:"user_id"`
	TransactionId string           `json:"transaction_id"`
	State         TransactionState `json:"state"`
	Amount        string           `json:"amount"`
}

type UserTransactionRepository interface {
	IsExistingUserTransaction(ctx context.Context, transactionId string) (bool, error)
	AddTransaction(ctx context.Context, userTransaction UserTransaction) error
	GetAllTransactionsByUserId(ctx context.Context, userId uint64) ([]UserTransaction, error)
}

type userTransactionRepository struct {
	db *pgxpool.Pool
}

func NewUserTransactionRepository(db *pgxpool.Pool) UserTransactionRepository {
	return &userTransactionRepository{db: db}
}

func (r *userTransactionRepository) IsExistingUserTransaction(ctx context.Context, transactionId string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM transactions WHERE transaction_id = $1)`

	var exists bool
	err := r.db.QueryRow(ctx, query, transactionId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *userTransactionRepository) AddTransaction(ctx context.Context, userTransaction UserTransaction) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO transactions (user_id, transaction_id, state, amount)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (transaction_id) DO NOTHING
	`

	tag, err := tx.Exec(ctx, query, userTransaction.UserId, userTransaction.TransactionId, userTransaction.State, userTransaction.Amount)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("transaction already exists: %s", userTransaction.TransactionId)
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (r *userTransactionRepository) GetAllTransactionsByUserId(ctx context.Context, userId uint64) ([]UserTransaction, error) {
	query := `
		SELECT user_id, transaction_id, state, amount
		FROM transactions
		WHERE user_id = $1
	`

	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []UserTransaction

	for rows.Next() {
		var t UserTransaction
		err := rows.Scan(&t.UserId, &t.TransactionId, &t.State, &t.Amount)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
