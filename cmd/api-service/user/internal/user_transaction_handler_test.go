package internal_test

import (
	"context"
	"entain-golang-task/cmd/api-service/user/internal"
	"entain-golang-task/pkg"
	pkgMock "entain-golang-task/pkg/mocks"
	"entain-golang-task/pkg/utils"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserTransactionHandler_SaveUserTransaction(t *testing.T) {
	t.Run("negative balance", func(t *testing.T) {
		userTransactionRepository := new(pkgMock.UserTransactionRepository)
		userTransactionRepository.On("IsExistingUserTransaction", context.TODO(), "transaction-1").Return(false, nil)
		userTransactions := []pkg.UserTransaction{{
			UserId:        1,
			TransactionId: "transaction-1",
			State:         "win",
			Amount:        "20.0",
		}}

		userTransactionRepository.On("GetAllTransactionsByUserId", context.TODO(), uint64(1)).Return(userTransactions, nil)

		handler := internal.NewUserTransactionHandlerWithInterfaces(zerolog.Logger{}, userTransactionRepository)

		transactionOutput, err := handler.SaveUserTransaction(context.TODO(), uint64(1), internal.UserTransactionInput{
			State:         "lose",
			Amount:        "21.3",
			TransactionId: "transaction-1",
		})

		assert.Nil(t, transactionOutput)
		assert.Exactly(t, utils.ErrAccountBalanceCannotBeNegative, err)
	})
}