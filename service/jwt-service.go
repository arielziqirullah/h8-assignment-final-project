package service

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateToken(userID uint, email string) (string, int64, error)
	ValidateToken(token string) (interface{}, error)
}

type jwtService struct {
	secretKey string
}

func NewJwtService() JwtService {
	return &jwtService{
		secretKey: os.Getenv("SECRET_KEY"),
	}
}

func (s *jwtService) GenerateToken(userID uint, email string) (string, int64, error) {

	expired := time.Now().Add(time.Hour * 72).Unix()

	claims := jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    expired,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", 0, err
	}

	return signedToken, expired, nil
}

func (s *jwtService) ValidateToken(token string) (interface{}, error) {
	bearer := strings.HasPrefix(token, "Bearer")

	if !bearer {
		return nil, errors.New("invalid token")
	}

	stringToken := strings.Split(token, " ")[1]

	getToken, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid parse token")
		}
		return []byte(s.secretKey), nil
	})

	if _, ok := getToken.Claims.(jwt.MapClaims); !ok && !getToken.Valid {
		return nil, errors.New("invalid get token")
	}

	return getToken.Claims.(jwt.MapClaims), nil
}
