package dto

type CustomerResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Pincode     string `json:"pincode"`
	DateOfBirth string `json:"dob"`
	Status      string `json:"status"`
}
