package controller

import (
	"h8-assignment-final-project/helpers"
	"h8-assignment-final-project/models"
	"h8-assignment-final-project/service"
	"h8-assignment-final-project/web/response"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PhotoController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	InsertPhoto(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

type photoController struct {
	photoService service.PhotoService
}

func NewPhotoController(photoService service.PhotoService) PhotoController {
	return &photoController{
		photoService: photoService,
	}
}

func (c *photoController) FindAll(ctx *gin.Context) {

	getAll := c.photoService.FindAll()

	if getAll == nil {
		response := response.BuildErrorResponse("Failed to get all photos", http.StatusNotFound, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, getAll)
	ctx.JSON(http.StatusOK, response)

}

func (c *photoController) FindByID(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		response := response.BuildErrorResponse("Failed to get photo", http.StatusBadRequest, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	photoID := uint(id)

	photo, err := c.photoService.FindByID(photoID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to get photo", http.StatusNotFound, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, photo)
	ctx.JSON(http.StatusOK, response)
}

func (c *photoController) InsertPhoto(ctx *gin.Context) {

	var photo models.Photo
	contentType := ctx.Request.Header.Get("Content-Type")
	if contentType != "application/json" {
		errorBind := ctx.ShouldBindJSON(&photo)
		if errorBind != nil {
			log.Println("Error binding JSON", errorBind)
			response := response.BuildErrorResponse("Failed to bind json", http.StatusBadRequest, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	} else {
		errorBind := ctx.ShouldBind(&photo)
		if errorBind != nil {
			log.Println("Error binding JSON", errorBind)
			response := response.BuildErrorResponse("Failed to bind", http.StatusBadRequest, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["userID"].(float64))
	photo.User_id = userID

	photo, err := c.photoService.InsertPhoto(photo)
	if err != nil {
		response := response.BuildErrorResponse(err.Error(), http.StatusBadRequest, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, photo)
	ctx.JSON(http.StatusOK, response)
}

func (c *photoController) UpdatePhoto(ctx *gin.Context) {

	var photo models.Photo
	contentType := ctx.Request.Header.Get("Content-Type")
	if contentType != "application/json" {
		errorBind := ctx.ShouldBindJSON(&photo)
		if errorBind != nil {
			log.Println("Error binding JSON", errorBind)
			response := response.BuildErrorResponse("Failed to bind json", http.StatusBadRequest, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	} else {
		errorBind := ctx.ShouldBind(&photo)
		if errorBind != nil {
			log.Println("Error binding", errorBind)
			response := response.BuildErrorResponse("Failed to bind", http.StatusBadRequest, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	// get photo id
	getPhotoData := ctx.MustGet("photoData").(*models.Photo)
	photo.ID = getPhotoData.ID

	// get user id
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["userID"].(float64))
	photo.User_id = userID

	updateData, err := c.photoService.UpdatePhoto(photo)
	if err != nil {
		log.Println("Error updating photo", err)
		response := response.BuildErrorResponse("Failed to update photo", http.StatusBadRequest, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, updateData)
	ctx.JSON(http.StatusOK, response)
}

func (c *photoController) DeletePhoto(ctx *gin.Context) {

	// get photo id
	getPhotoData := ctx.MustGet("photoData").(*models.Photo)

	err := c.photoService.DeletePhoto(getPhotoData.ID)
	if err != nil {
		log.Println("Error deleting photo", err)
		response := response.BuildErrorResponse("Failed to delete photo", http.StatusBadRequest, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, helpers.EmptyObj{})
	ctx.JSON(http.StatusOK, response)

}
