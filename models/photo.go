package models

import (
	"h8-assignment-final-project/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title     string    `json:"title" valid:"required~title is required"`
	Caption   string    `json:"caption"`
	Photo_url string    `json:"photo_url" valid:"required~photo_url is required"`
	User_id   uint      `gorm:"not null" json:"user_id"`
	Comment   []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Photo_id" json:"comment"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		helpers.LogIfError(errCreate, "Error validating photo")
		return errCreate
	}

	return
}
