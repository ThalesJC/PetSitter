package models

import (
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	User      User
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	Weight    string `json:"weight"`
	Size      string `json:"size"`
	Species   string `json:"species"`
	Gender    string `json:"gender"`
	Neutered  bool   `json:"neutered"`
	CoatColor string `json:"coat_color"`
	Picture   string `json:"picture"`
}
