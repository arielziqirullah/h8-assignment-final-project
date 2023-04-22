package models

import (
	"h8-assignment-final-project/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name             string `json:"name" valid:"required~name is required"`
	Social_media_url string `json:"social_media_url" valid:"required~social_media_url is required"`
	User_id          uint   `gorm:"not null" json:"user_id"`
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)
	if errCreate != nil {
		helpers.LogIfError(errCreate, "Error validating social media")
		return errCreate
	}

	return
}
