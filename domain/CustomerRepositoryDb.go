package domain

import (
	"database/sql"
	"github.com/alfiankan/gorest/errs"
	"github.com/alfiankan/gorest/logger"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	findByIdSql := "SELECT * FROM customers WHERE customer_id=?"
	result := d.client.QueryRow(findByIdSql, id)
	var c Customer
	err := result.Scan(&c.id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Data Not Found")
		} else {
			logger.Error("Error Query")
			return nil, errs.NewUnexpectedError("Unexpected Database problem")
		}

	}
	return &c, nil
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, error) {
	var result *sql.Rows
	var err error
	if status == "" {

		findAllSql := "SELECT * FROM customers"
		result, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "SELECT * FROM customers WHERE status = ?"
		result, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error Query " + err.Error())
		return nil, err
	}
	customers := make([]Customer, 0)
	for result.Next() {
		var c Customer
		err := result.Scan(&c.id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			logger.Error("Error Scanning " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

//helper db
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:12345@(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
