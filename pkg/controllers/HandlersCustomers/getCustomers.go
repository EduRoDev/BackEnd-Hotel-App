package handlerscustomers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"
	implements "github.com/DxKaizer/Hotel_Backend/pkg/services/Implements"
	"github.com/gorilla/mux"
)

type Customer struct {
	l  *log.Logger
	cr *implements.CustomersService
}

func NewLogger(l *log.Logger) *Customer {
	return &Customer{l, &implements.CustomersService{}}
}

func (c *Customer) GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	customers := c.cr.GetAllCustomers()
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(customers)
}

func (c *Customer) GetCustomersID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := map[string]string{
			"Status":  "Error",
			"Message": "Invalid customer ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}
	customer := entities.Customer{Id: idStr}
	Findcustomer := c.cr.GetCustomersID(customer)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Findcustomer)

}
