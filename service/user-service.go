package service

import (
	"h8-assignment-final-project/helpers"
	"h8-assignment-final-project/models"
	"h8-assignment-final-project/repository"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	InsertUser(user *models.User) (*models.User, error)
	IsDuplicateEmail(email string) (bool, error)
	VerifyCredential(username string, password string) (interface{}, error)
	MinAge(age int) bool
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) InsertUser(user *models.User) (*models.User, error) {
	userCreate := models.User{}
	err := smapping.FillStruct(&userCreate, smapping.MapFields(&user))
	helpers.LogIfError(err, "Error mapping user")

	process, err := s.userRepository.InsertUser(&userCreate)
	if err != nil {
		helpers.LogIfError(err, "Error inserting user")
		return &models.User{}, err
	}

	return process, nil
}

func (s *userService) VerifyCredential(username string, password string) (interface{}, error) {
	user, err := s.userRepository.VerifyCredential(username, password)
	if err != nil {
		return nil, err
	}

	if value, ok := user.(models.User); ok {
		comparedPassword := comparedPassword(value.Password, []byte(password))
		if value.Username == username && comparedPassword {
			return user, nil
		}
		return false, nil
	}
	return false, nil
}

func comparedPassword(hashPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (s *userService) IsDuplicateEmail(email string) (bool, error) {

	bool := false
	process, err := s.userRepository.IsDuplicateEmail(email)

	if err != nil {
		return false, err
	}

	if process > 0 {
		bool = true
		return bool, nil
	}
	return bool, nil
}

func (s *userService) MinAge(age int) bool {
	return age >= 8
}
