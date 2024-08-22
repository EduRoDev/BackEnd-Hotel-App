package model

type Hotel struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stars       int    `json:"stars"`
	Rooms       int    `json:"rooms"`
}

type Room struct {
	Id          int    `json:"id"`
	HotelId     int    `json:"hotel_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
