package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "SELECT * FROM customers"
	query, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error Query", err.Error())
		return nil, err
	}
	customers := make([]Customer, 0)
	for query.Next() {
		var c Customer
		err := query.Scan(&c.id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error Scanning", err.Error())
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
