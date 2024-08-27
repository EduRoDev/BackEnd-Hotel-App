package interfaces

import entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"

type CustomersInterface interface {
	GetAllCustomers() []entities.Customer
	GetCustomersID(customer entities.Customer) entities.Customer
	CreateCustomers(customer entities.Customer) map[string]interface{}
	EditCustomers(customer entities.Customer) map[string]interface{}
	DeleteCustomers(customer entities.Customer) map[string]interface{}
}
