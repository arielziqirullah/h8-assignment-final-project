package repository

import (
	"h8-assignment-final-project/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	FindAll() ([]models.SocialMedia, error)
	FindByID(id uint) (models.SocialMedia, error)
	InsertSocialMedia(socialMedia models.SocialMedia) (models.SocialMedia, error)
	UpdateSocialMedia(socialMedia models.SocialMedia) (models.SocialMedia, error)
	DeleteSocialMedia(id uint) error
}

type socialMediaConnection struct {
	connection *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaConnection{
		connection: db,
	}
}

func (db *socialMediaConnection) FindAll() ([]models.SocialMedia, error) {
	var socialMedia []models.SocialMedia
	err := db.connection.Find(&socialMedia).Error
	if err != nil {
		return []models.SocialMedia{}, err
	}
	return socialMedia, nil
}

func (db *socialMediaConnection) FindByID(id uint) (models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := db.connection.Where("id = ?", id).Take(&socialMedia).Error
	if err != nil {
		return models.SocialMedia{}, err
	}
	return socialMedia, nil
}

func (db *socialMediaConnection) InsertSocialMedia(socialMedia models.SocialMedia) (models.SocialMedia, error) {
	err := db.connection.Create(&socialMedia).Error
	if err != nil {
		return models.SocialMedia{}, err
	}
	return socialMedia, nil
}

func (db *socialMediaConnection) UpdateSocialMedia(socialMedia models.SocialMedia) (models.SocialMedia, error) {
	err := db.connection.Save(&socialMedia).Error
	if err != nil {
		return models.SocialMedia{}, err
	}
	return socialMedia, nil
}

func (db *socialMediaConnection) DeleteSocialMedia(id uint) error {
	err := db.connection.Delete(&models.SocialMedia{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
