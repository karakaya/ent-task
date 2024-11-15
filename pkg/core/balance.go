package core

import (
	"entain-golang-task/pkg/utils"
	"math/big"
)

func ValidateAmountBigRat(amountStr string) (string, error) {
	amount := new(big.Rat)
	_, ok := amount.SetString(amountStr)
	if !ok {
		return "", utils.ErrInvalidAmount
		//panic("invalid amount") could panic be more appropriate?
	}

	return amount.FloatString(2), nil
}

func SumTransactions(transactions []string) string {
	totalBalance := new(big.Rat)

	for _, amountStr := range transactions {
		amountRat := new(big.Rat)
		amountRat.SetString(amountStr)
		totalBalance.Add(totalBalance, amountRat)
	}

	balanceStr := totalBalance.FloatString(2)

	//additional validation for (-) negative balance?
	return balanceStr
}
