package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"
	implements "github.com/DxKaizer/Hotel_Backend/pkg/services/Implements"
	"github.com/DxKaizer/Hotel_Backend/pkg/services/interfaces"
	"github.com/gorilla/mux"
)

type Hotel struct {
	l  *log.Logger
	hs interfaces.HotelInterface
}

func NewLogger(l *log.Logger) *Hotel {
	return &Hotel{l, &implements.HotelService{}}
}

func (h *Hotel) GetHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	hotels := h.hs.GetAllHotels()
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(hotels)
}

func (h *Hotel) GetHotelID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := map[string]string{
			"Status":  "Error",
			"Message": "Invalid hotel ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}
	hotel := entities.Hotel{Id: idStr}
	Findhotel := h.hs.GetHotelID(hotel)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Findhotel)
}
