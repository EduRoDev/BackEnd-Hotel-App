package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DxKaizer/Hotel_Backend/pkg/data"
	"github.com/DxKaizer/Hotel_Backend/pkg/database"
	"github.com/DxKaizer/Hotel_Backend/pkg/model"
)

func (h *Hotel) PostHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var hotel data.HotelDto
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		rp := map[string]string{
			"Status":  "Error",
			"Message": "Error in create a product",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	data := model.Hotel{Name: hotel.Name, Description: hotel.Description, Rooms: hotel.Rooms}
	database.Database.Save(&data)
	rp := map[string]string{
		"Status":  "¡Success!",
		"Message": "Hotel created successfully",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rp)
}
