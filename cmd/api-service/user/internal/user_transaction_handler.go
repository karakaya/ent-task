package internal

import (
	"context"
	"ent-golang-task/database"
	"ent-golang-task/pkg/core"
	"ent-golang-task/pkg/repository"
	"ent-golang-task/pkg/utils"

	"github.com/rs/zerolog"
)

type UserTransactionHandler struct {
	logger                    zerolog.Logger
	userTransactionRepository repository.UserTransactionRepository
}

func NewUserTransactionHandler(logger zerolog.Logger) *UserTransactionHandler {
	return NewUserTransactionHandlerWithInterfaces(logger, repository.NewUserTransactionRepository(database.DB))
}

func NewUserTransactionHandlerWithInterfaces(logger zerolog.Logger, userTransactionRepository repository.UserTransactionRepository) *UserTransactionHandler {
	return &UserTransactionHandler{
		logger:                    logger,
		userTransactionRepository: userTransactionRepository,
	}
}

func (h *UserTransactionHandler) SaveUserTransaction(ctx context.Context, userId uint64, input UserTransactionInput) (*UserTransactionOutput, error) {
	//TODO: do we need to check if the user exists?

	isExistingTransaction, err := h.userTransactionRepository.IsExistingUserTransaction(ctx, input.TransactionId)
	if err != nil {
		return nil, err
	}

	if isExistingTransaction {
		return nil, utils.ErrTransactionExists
	}

	userTransactions, err := h.userTransactionRepository.GetAllTransactionsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	userAccountBalance, err := core.SumAllTransactions(userTransactions)
	if err != nil {
		return nil, err
	}

	canAddTransaction, currentAccountBalance, err := core.CanAddTransaction(userAccountBalance, input.Amount, input.State)
	if err != nil {
		return nil, err
	}

	if !canAddTransaction {
		return nil, utils.ErrAccountBalanceCannotBeNegative
	}

	err = h.userTransactionRepository.AddTransaction(ctx, repository.UserTransaction{
		UserId:        userId,
		TransactionId: input.TransactionId,
		State:         input.State,
		Amount:        input.Amount,
	})

	if err != nil {
		return nil, err
	}

	output := &UserTransactionOutput{
		TransactionId:  input.TransactionId,
		AccountBalance: currentAccountBalance,
	}

	return output, err
}

type UserTransactionInput struct {
	State         repository.TransactionState `json:"state"`
	Amount        string                      `json:"amount"`
	TransactionId string                      `json:"transactionId"`
}

type UserTransactionOutput struct {
	TransactionId  string `json:"transactionId"`
	AccountBalance string `json:"accountBalance"`
}
