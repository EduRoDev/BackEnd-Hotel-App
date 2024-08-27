package handlersroom

import (
	"encoding/json"
	"net/http"
	"strconv"

	entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"
	"github.com/gorilla/mux"
)

func (rm *Room) DeleteRoom(w http.ResponseWriter, r *http.Request) {
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
	deletedRoom := rm.rs.DeleteRooms(room)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedRoom)
}
