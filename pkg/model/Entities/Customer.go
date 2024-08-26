package entities

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