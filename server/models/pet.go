package models

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Birthdate time.Time `json:"birthdate"`
	Breed     string    `json:"breed"`
	CoatColor string    `json:"coat_color"`
	Name      string    `json:"name"`
	Neutered  bool      `json:"neutered"`
	PhotoUrl  string    `json:"photo_url"`
	Sex       string    `json:"sex"`
	Size      byte      `json:"size"`
	Species   string    `json:"species"`
	UserID    uint      `json:"user_id"`
	Weight    float32   `json:"weight"`
}
