package handlersroom

import (
	"encoding/json"
	"net/http"
	"strconv"

	entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"
	"github.com/DxKaizer/Hotel_Backend/pkg/model/dto"
	"github.com/gorilla/mux"
)

func (rm *Room) PutRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := map[string]string{
			"Status":  "Error",
			"Message": "Invalid room ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	var roomDto dto.RoomsDTO
	if err := json.NewDecoder(r.Body).Decode(&roomDto); err != nil {
		rp := map[string]string{
			"Status":  "Error",
			"Message": "Error decoding request body",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	room := entities.Room{Id: idStr, Name: roomDto.Name, Description: roomDto.Description, Price: roomDto.Price, HotelId: roomDto.HotelId}
	updatedRoom := rm.rs.EditRooms(room)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedRoom)
}
