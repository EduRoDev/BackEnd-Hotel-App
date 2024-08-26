package entities

type Room struct {
	Id          int    `json:"id"`
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Description string `gorm:"type:varchar(100)" json:"description"`
	Price       int    `json:"price"`
	HotelId     int    `json:"hotel_id"`
	Hotel       Hotel  `gorm:"foreignKey:HotelId" json:"hotel"`
}

type Rooms []Room