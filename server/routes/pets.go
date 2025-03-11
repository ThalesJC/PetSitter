package routes

type Pet struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	TutorId   string `json:"tutor_id" gorm:"foreignKey:ID"`
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
