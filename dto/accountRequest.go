package dto

import "github.com/71anshuman/banking-go/errs"

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("Amount should be minimum 5000")
	}

	if r.AccountType != "saving" && r.AccountType != "checking" {
		return errs.NewValidationError("Invalid account type")
	}

	return nil
}
