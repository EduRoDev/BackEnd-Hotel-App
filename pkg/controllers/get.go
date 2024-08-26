package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	implements "github.com/DxKaizer/Hotel_Backend/pkg/services/Implements"
	"github.com/DxKaizer/Hotel_Backend/pkg/services/interfaces"
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
