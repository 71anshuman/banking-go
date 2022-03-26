package domain

import (
	"github.com/71anshuman/banking-go/dto"
	"github.com/71anshuman/banking-go/errs"
)

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	TransactionType string  `db:"transaction_type"`
	Amount          float64 `db:"amount"`
	TransactionDate string  `db:"transaction_date"`
}

type TransactionRepository interface {
	Save(Transaction) (*Transaction, *errs.AppError)
}

func (t Transaction) IsWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
