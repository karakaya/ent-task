package core

import (
	"entain-golang-task/pkg"
	"entain-golang-task/pkg/utils"
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

func SumOfTransactions(transactions []string, transactionState pkg.TransactionState) (string, error) {
	totalBalance := new(big.Rat)

	for _, amountStr := range transactions {
		amountRat := new(big.Rat)
		amountRat.SetString(amountStr)
		totalBalance.Add(totalBalance, amountRat)

		switch transactionState {
		case pkg.StateWin:
			totalBalance.Add(totalBalance, amountRat)
		case pkg.StateLose:
			totalBalance.Sub(totalBalance, amountRat)
		default:
			return "", utils.ErrInvalidState
		}
	}

	balanceStr := totalBalance.FloatString(2)

	//additional validation for (-) negative balance?
	if totalBalance.Sign() < 0 {
		return "", fmt.Errorf("negative balance is not allowed")
	}

	return balanceStr, nil
}

func CanAddTransaction(currentBalanceStr string, amountStr string, transactionState pkg.TransactionState) (bool, error) {
	currentBalance := new(big.Rat)
	_, ok := currentBalance.SetString(currentBalanceStr)
	if !ok {
		return false, fmt.Errorf("invalid current balance format: %s", currentBalanceStr)
	}

	amountRat := new(big.Rat)
	_, ok = amountRat.SetString(amountStr)
	if !ok {
		return false, fmt.Errorf("invalid amount format: %s", amountStr)
	}

	resultBalance := new(big.Rat).Set(currentBalance)

	switch transactionState {
	case pkg.StateWin:
		resultBalance.Add(resultBalance, amountRat)
	case pkg.StateLose:
		resultBalance.Sub(resultBalance, amountRat)
	default:
		return false, utils.ErrInvalidState
	}

	if resultBalance.Sign() < 0 {
		return false, fmt.Errorf("transaction would result in negative balance")
	}

	return true, nil
}