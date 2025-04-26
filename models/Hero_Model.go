package model

import (
	"github.com/google/uuid"
)



type HeroModel struct {
	ID uuid.UUID `json:"id"`
	Surname string `json:"surname"`
	Profession string `json:"profession"`
	Description string `json:"description"`
	Location string `json:"location"`
	IsAvailable bool `json:"is_available"`
	ImageUrl string `json:"image_url"`
}