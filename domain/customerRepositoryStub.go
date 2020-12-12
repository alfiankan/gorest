package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	Customers := []Customer{
		{"1001", "alfiankan", "Solo", "57781", "1999-12-11", "1"},
		{"1002", "soni", "Solo", "57781", "1999-12-11", "1"},
	}
	return CustomerRepositoryStub{Customers}
}
