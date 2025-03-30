package models

import "time"

type Pet struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	Tutor     User      `json:"tutor"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
	Weight    float32   `json:"weight"`
	Size      byte      `json:"size"`
	Species   string    `json:"species"`
	Gender    string    `json:"gender"`
	Neutered  bool      `json:"neutered"`
	CoatColor string    `json:"coat_color"`
	Picture   string    `json:"picture"`
}
