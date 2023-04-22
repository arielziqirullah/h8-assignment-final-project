package repository

import (
	"h8-assignment-final-project/models"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user *models.User) (*models.User, error)
	VerifyCredential(username string, password string) (interface{}, error)
	IsDuplicateEmail(email string) (int, error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user *models.User) (*models.User, error) {
	err := db.connection.Create(user).Error
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (db *userConnection) VerifyCredential(username string, password string) (interface{}, error) {
	var user models.User
	err := db.connection.Where("username = ?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *userConnection) IsDuplicateEmail(email string) (int, error) {
	var user models.User
	err := db.connection.Where("email = ?", email).Take(&user).Error
	if err != nil {
		log.Println(err)
	}

	return len(user.Email), nil
}
