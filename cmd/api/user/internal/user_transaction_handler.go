package internal

import (
	"context"
	"entain-golang-task/database"
	"entain-golang-task/pkg"
	"entain-golang-task/pkg/core"
	"entain-golang-task/pkg/utils"

	"github.com/rs/zerolog"
)

type UserTransactionHandler struct {
	logger                    zerolog.Logger
	userRepository            pkg.UserRepository
	userTransactionRepository pkg.UserTransactionRepository
}

func NewUserHandler(logger zerolog.Logger) *UserTransactionHandler {
	return &UserTransactionHandler{
		logger:                    logger,
		userRepository:            pkg.NewUserRepository(logger),
		userTransactionRepository: pkg.NewUserTransactionRepository(database.DB),
	}
}

func (h *UserTransactionHandler) Handle(ctx context.Context, userId uint64, input UserTransactionInput) error {
	isExistingTransaction, err := h.userTransactionRepository.IsExistingUserTransaction(ctx, input.TransactionId)
	if err != nil {
		return err
	}
	if isExistingTransaction {
		return utils.ErrTransactionExists
	}

	userTransactions, err := h.userTransactionRepository.GetAllTransactionsByUserId(ctx, userId)
	if err != nil {
		return err
	}

	userAccountBalance, err := core.SumAllTransactions(userTransactions)
	if err != nil {
		return err
	}

	//TODO: current account balance will be added to the response
	canAddTransaction, _, err := core.CanAddTransaction(userAccountBalance, input.Amount, input.State)
	if err != nil {
		return err
	}

	if !canAddTransaction {
		return utils.ErrAccountBalanceCannotBeNegative
	}

	err = h.userTransactionRepository.AddTransaction(ctx, pkg.UserTransaction{
		UserId:        userId,
		TransactionId: input.TransactionId,
		State:         input.State,
		Amount:        input.Amount, //TODO: will address to balance check
	})

	return err
}

type UserTransactionInput struct {
	State         pkg.TransactionState `json:"state"`
	Amount        string               `json:"amount"`
	TransactionId string               `json:"transactionId"`
}
