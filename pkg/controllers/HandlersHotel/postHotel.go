package controllers

import (
	"encoding/json"
	"net/http"

	entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"
	"github.com/DxKaizer/Hotel_Backend/pkg/model/dto"
)

func (h *Hotel) PostHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var hotel dto.HotelDto
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		rp := map[string]string{
			"Status":  "Error",
			"Message": "Error in create a product",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	data := entities.Hotel{Name: hotel.Name, Description: hotel.Description, Rooms: hotel.Rooms}
	
	createdHotel := h.hs.CreateHotels(data)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdHotel)
}
