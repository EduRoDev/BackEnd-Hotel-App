package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DxKaizer/Hotel_Backend/pkg/database"
	"github.com/DxKaizer/Hotel_Backend/pkg/model"
)

type Hotel struct {
	l *log.Logger
}

func NewLogger(l *log.Logger) *Hotel {
	return &Hotel{l}
}

func (h *Hotel) GetHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data := model.Hotels{}
	database.Database.Find(&data)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(data)
}
