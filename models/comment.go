package models

import (
	"h8-assignment-final-project/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message  string `json:"message" valid:"required~message is required"`
	User_id  uint   `gorm:"not null" json:"user_id"`
	Photo_id uint   `gorm:"not null" json:"photo_id"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		helpers.LogIfError(errCreate, "Error validating comment")
		return errCreate
	}

	return
}
