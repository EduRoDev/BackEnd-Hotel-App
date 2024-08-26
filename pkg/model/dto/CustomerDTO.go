package dto

type CustomerDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	RoomID   int    `json:"room_id"`
}
