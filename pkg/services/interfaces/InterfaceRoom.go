package interfaces

import entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"

type RoomInterface interface {
	GetAllRooms() []entities.Room
	GetRoomID(room entities.Room) entities.Room
	CreateRooms(room entities.Room) map[string]interface{} // Cambiado aquí
	EditRooms(room entities.Room) entities.Room
	DeleteRooms(room entities.Room) entities.Room
}
