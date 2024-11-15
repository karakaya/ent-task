package internal_test

import (
	"context"
	"ent-golang-task/cmd/api-service/user/internal"
	"ent-golang-task/pkg"
	pkgMock "ent-golang-task/pkg/mocks"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserAccountBalanceHandler_GetAccountBalance(t *testing.T) {
	t.Run("get account balance 20.3", func(t *testing.T) {
		userTransactionRepository := new(pkgMock.UserTransactionRepository)
		userTransactions := []pkg.UserTransaction{{
			UserId:        1,
			TransactionId: "transaction-1",
			State:         "win",
			Amount:        "20.0",
		},
			{
				UserId:        1,
				TransactionId: "transaction-3",
				State:         "win",
				Amount:        "10",
			},
			{
				UserId:        1,
				TransactionId: "transaction-3",
				State:         "lose",
				Amount:        "10",
			},
			{
				UserId:        1,
				TransactionId: "transaction-3",
				State:         "win",
				Amount:        "0.3",
			},
		}

		userTransactionRepository.On("GetAllTransactionsByUserId", context.TODO(), uint64(1)).Return(userTransactions, nil)

		handler := internal.NewNewUserAccountBalanceHandlerWithInterfaces(zerolog.Logger{}, userTransactionRepository)

		expectedOutput := &internal.UserAccountBalanceOutput{
			UserId:  uint64(1),
			Balance: "20.30",
		}

		accountBalance, err := handler.GetAccountBalance(context.TODO(), uint64(1))
		assert.Nil(t, err)
		assert.Exactly(t, expectedOutput, accountBalance)
	})
}
