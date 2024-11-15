package core_test

import (
	"entain-golang-task/pkg"
	"entain-golang-task/pkg/core"
	"entain-golang-task/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateTransactionAmount(t *testing.T) {
	t.Run("valid amount", func(t *testing.T) {
		validatedAmount, err := core.ValidateTransactionAmount("15.11")
		assert.Nil(t, err)
		assert.Exactly(t, "15.11", validatedAmount)
	})

	t.Run("invalid amount - comma", func(t *testing.T) {
		validatedAmount, err := core.ValidateTransactionAmount("15,11")
		assert.NotNil(t, err)
		assert.Exactly(t, "", validatedAmount)
	})

	t.Run("invalid amount - letter", func(t *testing.T) {
		validatedAmount, err := core.ValidateTransactionAmount("a15,11")
		assert.NotNil(t, err)
		assert.Exactly(t, "", validatedAmount)
	})
}

func TestSumAllTransactions(t *testing.T) {
	t.Run("positive, 15.11", func(t *testing.T) {
		userTransactions := []pkg.UserTransaction{{
			UserId:        1,
			TransactionId: "transaction-1",
			State:         "win",
			Amount:        "10.0",
		},
			{
				UserId:        1,
				TransactionId: "transaction-2",
				State:         "win",
				Amount:        "10.11",
			},
			{
				UserId:        1,
				TransactionId: "transaction-3",
				State:         "lose",
				Amount:        "5.0",
			},
		}

		sumOfAllTransactions, err := core.SumAllTransactions(userTransactions)

		assert.Nil(t, err)
		assert.Exactly(t, "11.11", sumOfAllTransactions)
	})
}

func TestCanAddTransaction(t *testing.T) {
	t.Run("positive balance - can add transaction", func(t *testing.T) {
		canAddTransaction, newBalance, err := core.CanAddTransaction("10.10", "10", "lose")
		assert.Nil(t, err)
		assert.Exactly(t, "0.10", newBalance)
		assert.True(t, canAddTransaction)
	})

	t.Run("positive balance - can add transaction", func(t *testing.T) {
		canAddTransaction, newBalance, err := core.CanAddTransaction("10.10", "20", "lose")
		
		assert.Exactly(t, utils.ErrAccountBalanceCannotBeNegative, err)
		assert.Exactly(t, "", newBalance)
		assert.False(t, canAddTransaction)
	})
}
