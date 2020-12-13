package domain

import "github.com/alfiankan/gorest/errs"

type Customer struct {
	id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {

	// status ==1 status ==0 status == ""
	FindAll(string) ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
