package database

import (
	"fmt"

	entities "github.com/DxKaizer/Hotel_Backend/pkg/model/Entities"
)

func Migraciones() {
	if Database == nil {
		fmt.Print("Database connection is not established.\n")
		return
	}

	err := Database.AutoMigrate(entities.Hotel{}, entities.Room{}, entities.Customer{})
	if err != nil {
		fmt.Print("Error during migration:", err)
		return
	}

	fmt.Print("Database migration completed successfully.\n")
}
