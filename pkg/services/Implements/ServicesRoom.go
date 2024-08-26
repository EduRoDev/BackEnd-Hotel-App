package implements

import (
	"github.com/DxKaizer/Hotel_Backend/pkg/database"
	entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"
)

type RoomService struct{}

func (rs *RoomService) GetAllRooms() []entities.Room {
	var rooms []entities.Room
	result := database.Database.Preload("Hotel").Find(&rooms)
	if result.Error != nil {
		return nil
	}
	return rooms
}

func (rs *RoomService) GetRoomID(room entities.Room) entities.Room {
	result := database.Database.Preload("Hotel").First(&room, room.Id)
	if result.Error != nil {
		return entities.Room{}
	}
	return room
}

func (rs *RoomService) CreateRooms(room entities.Room) map[string]interface{} {
	result := database.Database.Create(&room)
	if result.Error != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Error al crear la habitación",
		}
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Habitación creada exitosamente",
	}
}

func (rs *RoomService) EditRooms(room entities.Room) entities.Room {
	result := database.Database.Model(&room).Updates(room)
	if result.Error != nil {
		return entities.Room{}
	}
	return room
}

func (rs *RoomService) DeleteRooms(room entities.Room) entities.Room {
	result := database.Database.Delete(&room)
	if result.Error != nil {
		return entities.Room{}
	}
	return room
}
