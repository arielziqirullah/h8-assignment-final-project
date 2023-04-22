package models

import (
	"h8-assignment-final-project/helpers"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string        `gorm:"uniqueIndex;not null;varchar(100)" json:"username" valid:"required~username is required"`
	Email       string        `gorm:"uniqueIndex;not null;varchar(100)" json:"email" valid:"email,required~email is required"`
	Password    string        `gorm:"not null;varchar(100)" json:"password" valid:"minstringlength(6),required~password is required"`
	Age         int           `json:"age" valid:"required~age is required"`
	Comment     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:User_id" json:"comment"`
	Photo       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:User_id" json:"photo"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:User_id" json:"social_media"`
}

func (u *User) BeforeCreate(txc *gorm.DB) (err error) {
	result, errCreate := govalidator.ValidateStruct(u)
	if !result {
		helpers.LogIfError(errCreate, "error validating user")
		return errCreate
	}

	u.Password = hashAndSalt([]byte(u.Password))

	return
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	helpers.LogIfError(err, "Error hashing password")

	return string(hash)
}
