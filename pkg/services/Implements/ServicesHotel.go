package implements

import (
	"github.com/DxKaizer/Hotel_Backend/pkg/database"
	entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"
)

type HotelService struct{}

func (hs *HotelService) GetAllHotels() []entities.Hotel {
	var hotels []entities.Hotel
	result := database.Database.Find(&hotels)
	if result.Error != nil {
		return nil
	}
	return hotels
}

func (hs *HotelService) CreateHotels(hotel entities.Hotel) entities.Hotel {
	result := database.Database.Create(&hotel)
	if result.Error != nil {
		return entities.Hotel{}
	}
	return hotel
}

func (hs *HotelService) EditHotels(hotel entities.Hotel) entities.Hotel {
	result := database.Database.Model(&hotel).Updates(hotel)
	if result.Error != nil {
		return entities.Hotel{}
	}
	return hotel
}
