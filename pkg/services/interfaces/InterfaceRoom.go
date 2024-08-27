package interfaces

import entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"

type RoomInterface interface {
	GetAllRooms() []entities.Room
	GetRoomID(room entities.Room) entities.Room
	CreateRooms(room entities.Room) map[string]interface{}
	EditRooms(room entities.Room) map[string]interface{}
	DeleteRooms(room entities.Room) map[string]interface{}
}
