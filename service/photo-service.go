package service

import (
	"h8-assignment-final-project/models"
	"h8-assignment-final-project/repository"
)

type PhotoService interface {
	FindAll() []models.Photo
	FindByID(id uint) (models.Photo, error)
	InsertPhoto(photo models.Photo) (models.Photo, error)
	UpdatePhoto(photo models.Photo) (models.Photo, error)
	DeletePhoto(id uint) error
}

type photoService struct {
	photoRepository repository.PhotoRepository
}

func NewPhotoService(photoRepo repository.PhotoRepository) PhotoService {
	return &photoService{
		photoRepository: photoRepo,
	}
}

func (s *photoService) FindAll() []models.Photo {
	photos, _ := s.photoRepository.FindAll()
	return photos
}

func (s *photoService) FindByID(id uint) (models.Photo, error) {
	photo, err := s.photoRepository.FindByID(id)
	if err != nil {
		return models.Photo{}, err
	}
	return photo, nil
}

func (s *photoService) InsertPhoto(photo models.Photo) (models.Photo, error) {
	newPhoto, err := s.photoRepository.InsertPhoto(photo)
	if err != nil {
		return models.Photo{}, err
	}
	return newPhoto, nil
}

func (s *photoService) UpdatePhoto(photo models.Photo) (models.Photo, error) {
	updatedPhoto, err := s.photoRepository.UpdatePhoto(photo)
	if err != nil {
		return models.Photo{}, err
	}
	return updatedPhoto, nil
}

func (s *photoService) DeletePhoto(id uint) error {
	err := s.photoRepository.DeletePhoto(id)
	if err != nil {
		return err
	}
	return nil
}
