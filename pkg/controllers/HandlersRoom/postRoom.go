package handlersroom

import (
	"encoding/json"
	"net/http"

	entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"
	"github.com/DxKaizer/Hotel_Backend/pkg/model/dto"
)

func (rm *Room) PostRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var room dto.RoomsDTO
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		rp := map[string]string{
			"Status":  "Error",
			"Message": "Error in create a product",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	data := entities.Room{Name: room.Name, Description: room.Description, Price: room.Price, HotelId: room.HotelId}
	createdRoom := rm.rs.CreateRooms(data)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdRoom)
}
