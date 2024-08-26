package entities

type Hotel struct {
	Id          int    `json:"id"`
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Description string `gorm:"type:varchar(100)" json:"description"`
	Rooms       int    `json:"rooms"`
}

type Hotels []Hotel