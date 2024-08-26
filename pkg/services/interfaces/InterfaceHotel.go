package interfaces

import entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"

type HotelInterface interface {
	GetAllHotels() []entities.Hotel
	CreateHotels(hotel entities.Hotel) entities.Hotel
	EditHotels(hotel entities.Hotel) entities.Hotel
	//DeleteHotels(hotel entities.Hotel) entities.Hotel
}
