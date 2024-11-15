package internal_test

import (
	"context"
	"entain-golang-task/cmd/api-service/user/internal"
	"entain-golang-task/pkg"
	pkgMock "entain-golang-task/pkg/mocks"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserAccountBalanceHandler_GetAccountBalance(t *testing.T) {
	t.Run("get account balance 20.3", func(t *testing.T) {
		transactionRepository := new(pkgMock.UserTransactionRepository)
		transactions := []pkg.UserTransaction{{
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

		transactionRepository.On("GetAllTransactionsByUserId", context.TODO(), uint64(1)).Return(transactions, nil)

		handler := internal.NewNewUserAccountBalanceHandlerWithInterfaces(zerolog.Logger{}, transactionRepository)

		expectedOutput := &internal.UserAccountBalanceOutput{
			UserId:  uint64(1),
			Balance: "20.30",
		}

		accountBalance, err := handler.GetAccountBalance(context.TODO(), uint64(1))
		assert.Nil(t, err)
		assert.Exactly(t, expectedOutput, accountBalance)
	})
}
