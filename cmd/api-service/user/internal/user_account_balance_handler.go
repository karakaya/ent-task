package internal

import (
	"context"
	"entain-golang-task/database"
	"entain-golang-task/pkg"
	"entain-golang-task/pkg/core"
	"github.com/rs/zerolog"
)

type UserAccountBalanceHandler struct {
	logger                    zerolog.Logger
	userTransactionRepository pkg.UserTransactionRepository
}

func NewUserAccountBalanceHandler(logger zerolog.Logger) *UserAccountBalanceHandler {
	return NewNewUserAccountBalanceHandlerWithInterfaces(logger, pkg.NewUserTransactionRepository(database.DB))
}

func NewNewUserAccountBalanceHandlerWithInterfaces(
	logger zerolog.Logger,
	userTransactionRepository pkg.UserTransactionRepository,
) *UserAccountBalanceHandler {
	return &UserAccountBalanceHandler{
		logger:                    logger,
		userTransactionRepository: userTransactionRepository,
	}
}

func (h *UserAccountBalanceHandler) GetAccountBalance(ctx context.Context, userId uint64) (*UserAccountBalanceOutput, error) {
	userTransactions, err := h.userTransactionRepository.GetAllTransactionsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	sumOfUserTransactions, err := core.SumAllTransactions(userTransactions)
	if err != nil {
		return nil, err
	}

	output := &UserAccountBalanceOutput{
		UserId:  userId,
		Balance: sumOfUserTransactions,
	}

	return output, nil
}

type UserAccountBalanceOutput struct {
	UserId  uint64 `json:"userId"`
	Balance string `json:"balance"`
}
