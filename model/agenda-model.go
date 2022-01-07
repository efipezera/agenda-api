package model

type Contact struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type Input struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
}
