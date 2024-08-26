package handlersroom

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

type Room struct {
	l  *log.Logger
	rs interfaces.RoomInterface
}

func NewLogger(l *log.Logger) *Room {
	return &Room{l, &implements.RoomService{}}
}

func (rm *Room) GetRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	rooms := rm.rs.GetAllRooms()
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(rooms)
}

func (rm *Room) GetRoomID(w http.ResponseWriter, r *http.Request) {
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
	room := entities.Room{Id: idStr}
	FindRoom := rm.rs.GetRoomID(room)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(FindRoom)
}
