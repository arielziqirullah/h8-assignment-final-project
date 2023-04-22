package service

import (
	"h8-assignment-final-project/models"
	"h8-assignment-final-project/repository"
)

type SocialMediaService interface {
	FindAll() []models.SocialMedia
	FindByID(id uint) (models.SocialMedia, error)
	InsertSocialMedia(socialMedia models.SocialMedia) (models.SocialMedia, error)
	UpdateSocialMedia(socialMedia models.SocialMedia) (models.SocialMedia, error)
	DeleteSocialMedia(id uint) error
}

type socialMediaService struct {
	socialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepo repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{
		socialMediaRepository: socialMediaRepo,
	}
}

func (s *socialMediaService) FindAll() []models.SocialMedia {
	socialMedia, _ := s.socialMediaRepository.FindAll()
	return socialMedia
}

func (s *socialMediaService) FindByID(id uint) (models.SocialMedia, error) {
	socialMedia, err := s.socialMediaRepository.FindByID(id)
	if err != nil {
		return models.SocialMedia{}, err
	}
	return socialMedia, nil
}

func (s *socialMediaService) InsertSocialMedia(socialMedia models.SocialMedia) (models.SocialMedia, error) {
	newSocialMedia, err := s.socialMediaRepository.InsertSocialMedia(socialMedia)
	if err != nil {
		return models.SocialMedia{}, err
	}
	return newSocialMedia, nil
}

func (s *socialMediaService) UpdateSocialMedia(socialMedia models.SocialMedia) (models.SocialMedia, error) {
	newSocialMedia, err := s.socialMediaRepository.UpdateSocialMedia(socialMedia)
	if err != nil {
		return models.SocialMedia{}, err
	}
	return newSocialMedia, nil
}

func (s *socialMediaService) DeleteSocialMedia(id uint) error {
	err := s.socialMediaRepository.DeleteSocialMedia(id)
	if err != nil {
		return err
	}
	return nil
}
