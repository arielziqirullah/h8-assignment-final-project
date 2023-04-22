package repository

import (
	"h8-assignment-final-project/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	FindAll() ([]models.Photo, error)
	FindByID(id uint) (models.Photo, error)
	InsertPhoto(photo models.Photo) (models.Photo, error)
	UpdatePhoto(photo models.Photo) (models.Photo, error)
	DeletePhoto(id uint) error
}

type photoConnection struct {
	connection *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoConnection{
		connection: db,
	}
}

func (db *photoConnection) FindAll() ([]models.Photo, error) {
	var photos []models.Photo
	err := db.connection.Model(&models.Photo{}).Preload("Comment").Find(&photos).Error
	if err != nil {
		return []models.Photo{}, err
	}
	return photos, nil
}

func (db *photoConnection) FindByID(id uint) (models.Photo, error) {
	var photo models.Photo
	err := db.connection.Where("id = ?", id).Take(&photo).Error
	if err != nil {
		return models.Photo{}, err
	}
	return photo, nil
}

func (db *photoConnection) InsertPhoto(photo models.Photo) (models.Photo, error) {
	err := db.connection.Create(&photo).Error
	if err != nil {
		return models.Photo{}, err
	}
	return photo, nil
}

func (db *photoConnection) UpdatePhoto(photo models.Photo) (models.Photo, error) {
	err := db.connection.Save(&photo).Error
	if err != nil {
		return models.Photo{}, err
	}
	return photo, nil
}

func (db *photoConnection) DeletePhoto(id uint) error {
	err := db.connection.Where("id = ?", id).Delete(&models.Photo{}).Error
	if err != nil {
		return err
	}
	return nil
}
