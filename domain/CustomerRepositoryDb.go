package domain

import (
	"database/sql"

	"github.com/71anshuman/banking-go/errs"
	"github.com/71anshuman/banking-go/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)

	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"

	switch status {
	case "active":
		findAllSql = findAllSql + " WHERE status = 1"
	case "inactive":
		findAllSql = findAllSql + " WHERE status = 0"
	}

	err := d.client.Select(&customers, findAllSql)
	if err != nil {
		logger.Info("Error while quering the customer table" + err.Error())
		return nil, errs.NewUnExpectedError(errs.DB_ERR)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	findByIdSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"
	var c Customer

	err := d.client.Get(&c, findByIdSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Info("Error while scanning customer " + err.Error())
			return nil, errs.NewUnExpectedError(errs.DB_ERR)
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
