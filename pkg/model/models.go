package model

import (
	"fmt"

	"github.com/DxKaizer/Hotel_Backend/pkg/database"
)

type Hotel struct {
	Id          int    `json:"id"`
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Description string `gorm:"type:varchar(100)" json:"description"`
	Rooms       int    `json:"rooms"`
}

type Hotels []Hotel

type Room struct {
	Id          int    `json:"id"`
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Description string `gorm:"type:varchar(100)" json:"description"`
	Price       int    `json:"price"`
	HotelId     int    `json:"hotel_id"`
	Hotel       Hotel  `gorm:"foreignKey:HotelId" json:"hotel"`
}

type Rooms []Room

type Customer struct {
	Id       int    `json:"id"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Email    string `gorm:"type:varchar(100)" json:"email"`
	Phone    string `json:"phone"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	RoomID   int    `json:"room_id"`
	Room     Room   `gorm:"foreignKey:RoomID" json:"room"`
}

type Customers []Customer

func Migraciones() {
	if database.Database == nil {
		fmt.Print("Database connection is not established.\n")
		return
	}

	err := database.Database.AutoMigrate(&Hotel{}, &Room{}, &Customer{})
	if err != nil {
		fmt.Print("Error during migration:", err)
		return
	}

	fmt.Print("Database migration completed successfully.\n")
}
