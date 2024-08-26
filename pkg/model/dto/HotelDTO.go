package dto

type HotelDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Rooms       int    `json:"rooms"`
}