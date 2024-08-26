package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"
	"github.com/gorilla/mux"
)

func (h *Hotel) DeleteHotel(w http.ResponseWriter, r *http.Request) {
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
	hotel := entities.Hotel{Id: id}
	deletedHotel := h.hs.DeleteHotels(hotel)
	if deletedHotel.Id == 0 {
		rp := map[string]string{
			"Status":  "Error",
			"Message": "Hotel not found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedHotel)
}
