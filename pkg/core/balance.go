package core

import (
	"ent-golang-task/pkg/repository"
	"ent-golang-task/pkg/utils"
	"fmt"
	"math/big"
)

func ValidateTransactionAmount(amountStr string) (string, error) {
	amount := new(big.Rat)
	_, ok := amount.SetString(amountStr)
	if !ok {
		return "", utils.ErrInvalidAmount
		//panic("invalid amount")
	}

	return amount.FloatString(2), nil
}

func SumAllTransactions(transactions []repository.UserTransaction) (string, error) {
	totalBalance := new(big.Rat)

	for _, transaction := range transactions {
		amountRat := new(big.Rat)
		amountRat.SetString(transaction.Amount)

		switch transaction.State {
		case repository.StateWin:
			totalBalance.Add(totalBalance, amountRat)
		case repository.StateLose:
			totalBalance.Sub(totalBalance, amountRat)
		default:
			return "", utils.ErrInvalidState
		}
	}

	balanceStr := totalBalance.FloatString(2)

	//additional validation for (-) negative balance?
	if totalBalance.Sign() < 0 {
		return "", utils.ErrAccountBalanceCannotBeNegative
	}

	return balanceStr, nil
}

func CanAddTransaction(currentBalanceStr string, amountStr string, transactionState repository.TransactionState) (bool, string, error) {
	currentBalance := new(big.Rat)
	_, ok := currentBalance.SetString(currentBalanceStr)
	if !ok {
		return false, "", fmt.Errorf("invalid current balance format: %s", currentBalanceStr)
	}

	amountRat := new(big.Rat)
	_, ok = amountRat.SetString(amountStr)
	if !ok {
		return false, "", fmt.Errorf("invalid amount format: %s", amountStr)
	}

	resultBalance := new(big.Rat).Set(currentBalance)

	switch transactionState {
	case repository.StateWin:
		resultBalance.Add(resultBalance, amountRat)
	case repository.StateLose:
		resultBalance.Sub(resultBalance, amountRat)
	default:
		return false, "", utils.ErrInvalidState
	}

	if resultBalance.Sign() < 0 {
		return false, "", utils.ErrAccountBalanceCannotBeNegative
	}

	currentAccountBalance := resultBalance.FloatString(2)

	return true, currentAccountBalance, nil
}
