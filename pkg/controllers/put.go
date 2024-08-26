package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"
	"github.com/DxKaizer/Hotel_Backend/pkg/model/dto"
	"github.com/gorilla/mux"
)

func (h *Hotel) PutHotels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vr := mux.Vars(r)
	idStr := vr["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		rp := map[string]string{
			"Status":  "Error",
			"Message": "Invalid hotel ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	var hotelDto dto.HotelDto
	if err := json.NewDecoder(r.Body).Decode(&hotelDto); err != nil {
		rp := map[string]string{
			"Status":  "Error",
			"Message": "Error decoding request body",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	hotel := entities.Hotel{Id: id, Name: hotelDto.Name, Description: hotelDto.Description, Rooms: hotelDto.Rooms}
	updatedHotel := h.hs.EditHotels(hotel)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedHotel)
}
