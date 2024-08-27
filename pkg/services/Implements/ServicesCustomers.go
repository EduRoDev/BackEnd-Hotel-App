package implements

import (
	"github.com/DxKaizer/Hotel_Backend/pkg/database"
	entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"
)

type CustomersService struct{}

func (cs *CustomersService) GetAllCustomers() []entities.Customer {
	var customers []entities.Customer
	result := database.Database.Preload("Room").Preload("Room.Hotel").Find(&customers)
	if result.Error != nil {
		return nil
	}
	return customers
}

func (cs *CustomersService) GetCustomersID(customer entities.Customer) entities.Customer {
	result := database.Database.Preload("Room").Preload("Room.Hotel").First(&customer, customer.Id)
	if result.Error != nil {
		return entities.Customer{}
	}
	return customer
}

func (cs *CustomersService) CreateCustomers(customer entities.Customer) map[string]interface{} {
	result := database.Database.Create(&customer)
	if result.Error != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Error al crear el cliente",
		}
	}
	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Cliente creado exitosamente",
	}
}

func (cs *CustomersService) EditCustomers(customer entities.Customer) map[string]interface{} {
	result := database.Database.Model(&customer).Updates(customer)
	if result.Error != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Error al editar el cliente",
		}
	}
	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Cliente editado exitosamente",
	}
}
