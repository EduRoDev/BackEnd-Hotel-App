package dto

type RoomsDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	HotelId     int    `json:"hotel_id"`
}