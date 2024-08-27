package handlerscustomers

import (
	"encoding/json"
	"net/http"

	entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"
	"github.com/DxKaizer/Hotel_Backend/pkg/model/dto"
)

func (c *Customer) PostCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var customer dto.CustomerDTO
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		rp := map[string]string{
			"Status":  "Error",
			"Message": "Invalid customer data",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}
	data := entities.Customer{Name: customer.Name, Email: customer.Email, Phone: customer.Phone, Password: customer.Password, RoomID: customer.RoomID}
	createdCustomer := c.cr.CreateCustomers(data)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCustomer)

}
