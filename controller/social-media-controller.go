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

type SocialMediaController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	InsertSocialMedia(ctx *gin.Context)
	UpdateSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

type socialMediaController struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaController(socialMediaService service.SocialMediaService) SocialMediaController {
	return &socialMediaController{
		socialMediaService: socialMediaService,
	}
}

func (c *socialMediaController) FindAll(ctx *gin.Context) {
	getAll := c.socialMediaService.FindAll()

	if getAll == nil {
		response := response.BuildErrorResponse("Failed to get all social media", http.StatusNotFound, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, getAll)
	ctx.JSON(http.StatusOK, response)

}

func (c *socialMediaController) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		response := response.BuildErrorResponse("Failed to process request", http.StatusBadRequest, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	socialMediaID := uint(id)

	photo, err := c.socialMediaService.FindByID(socialMediaID)
	if err != nil {
		log.Println(err)
		response := response.BuildErrorResponse("social media not found", http.StatusNotFound, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, photo)
	ctx.JSON(http.StatusOK, response)
}

func (c *socialMediaController) InsertSocialMedia(ctx *gin.Context) {

	var socialMedia models.SocialMedia
	contentType := ctx.Request.Header.Get("Content-Type")
	if contentType != "application/json" {
		errorBind := ctx.ShouldBindJSON(&socialMedia)
		if errorBind != nil {
			log.Println("Error binding JSON", errorBind)
			response := response.BuildErrorResponse("Failed to bind json", http.StatusBadRequest, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	} else {
		errorBind := ctx.ShouldBind(&socialMedia)
		if errorBind != nil {
			log.Println("Error binding JSON", errorBind)
			response := response.BuildErrorResponse("Failed to bind", http.StatusBadRequest, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["userID"].(float64))
	socialMedia.User_id = userID

	photo, err := c.socialMediaService.InsertSocialMedia(socialMedia)
	if err != nil {
		log.Println("Error inserting social media", err)
		response := response.BuildErrorResponse(err.Error(), http.StatusBadRequest, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, photo)
	ctx.JSON(http.StatusOK, response)
}

func (c *socialMediaController) UpdateSocialMedia(ctx *gin.Context) {

	var socialMedia models.SocialMedia
	contentType := ctx.Request.Header.Get("Content-Type")
	if contentType != "application/json" {
		errorBind := ctx.ShouldBindJSON(&socialMedia)
		if errorBind != nil {
			log.Println("Error binding JSON", errorBind)
			response := response.BuildErrorResponse("Failed to bind json", http.StatusBadRequest, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	} else {
		errorBind := ctx.ShouldBind(&socialMedia)
		if errorBind != nil {
			log.Println("Error binding JSON", errorBind)
			response := response.BuildErrorResponse("Failed to bind", http.StatusBadRequest, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	// get social media id
	getSocialMediaData := ctx.MustGet("socialMediaData").(*models.SocialMedia)
	socialMedia.ID = getSocialMediaData.ID

	// get user id
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["userID"].(float64))
	socialMedia.User_id = userID

	updateData, err := c.socialMediaService.UpdateSocialMedia(socialMedia)
	if err != nil {
		log.Println("Error updating social media", err)
		response := response.BuildErrorResponse("Failed to update social media", http.StatusBadRequest, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, updateData)
	ctx.JSON(http.StatusOK, response)

}

func (c *socialMediaController) DeleteSocialMedia(ctx *gin.Context) {

	getSocialMediaData := ctx.MustGet("socialMediaData").(*models.SocialMedia)

	err := c.socialMediaService.DeleteSocialMedia(getSocialMediaData.ID)
	if err != nil {
		log.Println("Error deleting social media", err)
		response := response.BuildErrorResponse("Failed to delete social media", http.StatusBadRequest, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, helpers.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
