package middlewares

import (
	"h8-assignment-final-project/helpers"
	"h8-assignment-final-project/models"
	"h8-assignment-final-project/web/response"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthorizationMiddleware interface {
	PhotoAuthorization() gin.HandlerFunc
	SocialMediaAuthorization() gin.HandlerFunc
	CommentAuthorization() gin.HandlerFunc
}

type authorizationMiddleware struct {
	connection *gorm.DB
}

func NewAuthorizationMiddleware(db *gorm.DB) AuthorizationMiddleware {
	return &authorizationMiddleware{
		connection: db,
	}
}

func (con *authorizationMiddleware) PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		photoID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(err)
			response := response.BuildErrorResponse("Failed to get photo", http.StatusBadRequest, helpers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		usersID := uint(userData["userID"].(float64))

		Photo := models.Photo{}

		err = con.connection.First(&Photo, uint(photoID)).Error
		if err != nil {
			log.Println(err)
			response := response.BuildErrorResponse("Data Not Found", http.StatusNotFound, helpers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusNotFound, response)
			return
		}

		if Photo.User_id != usersID {
			response := response.BuildErrorResponse("Unauthorized to access data", http.StatusUnauthorized, helpers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("photoData", &Photo)
		c.Next()
	}
}

func (con *authorizationMiddleware) SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		socialMediaID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(err)
			response := response.BuildErrorResponse("Failed to get social media", http.StatusBadRequest, helpers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		usersID := uint(userData["userID"].(float64))

		SocialMedia := models.SocialMedia{}

		err = con.connection.First(&SocialMedia, uint(socialMediaID)).Error
		if err != nil {
			log.Println(err)
			response := response.BuildErrorResponse("Data Not Found", http.StatusNotFound, helpers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusNotFound, response)
			return
		}

		if SocialMedia.User_id != usersID {
			response := response.BuildErrorResponse("Unauthorized to access data", http.StatusUnauthorized, helpers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("socialMediaData", &SocialMedia)
		c.Next()
	}
}

func (con *authorizationMiddleware) CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(err)
			response := response.BuildErrorResponse("Failed to get comment", http.StatusBadRequest, helpers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		usersID := uint(userData["userID"].(float64))

		Comment := models.Comment{}

		err = con.connection.First(&Comment, uint(commentID)).Error
		if err != nil {
			log.Println(err)
			response := response.BuildErrorResponse("Data Not Found", http.StatusNotFound, helpers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusNotFound, response)
			return
		}

		if Comment.User_id != usersID {
			response := response.BuildErrorResponse("Unauthorized to access data", http.StatusUnauthorized, helpers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("commentData", &Comment)
		c.Next()
	}
}
