package domain

import (
	"github.com/71anshuman/banking-go/dto"
	"github.com/71anshuman/banking-go/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Pincode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		DateOfBirth: c.DateOfBirth,
		City:        c.City,
		Pincode:     c.Pincode,
		Status:      c.GetStatusAsText(),
	}
}

func (c Customer) GetStatusAsText() string {
	status := "active"
	if c.Status == "0" {
		status = "inactive"
	}
	return status
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) ByID() ([]Customer, error) {
	return s.customers[0:1], nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "10001", Name: "Anshuman Lawania", City: "Noida", Pincode: "201301", DateOfBirth: "26-07-1997", Status: "1"},
		{Id: "10002", Name: "Archana Lawania", City: "Agra", Pincode: "282010", DateOfBirth: "19-02-1993", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
