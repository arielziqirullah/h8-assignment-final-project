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

type CommentController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	InsertComment(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type commentController struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) CommentController {
	return &commentController{
		commentService: commentService,
	}
}

func (c *commentController) FindAll(ctx *gin.Context) {

	getAll := c.commentService.FindAll()

	if getAll == nil {
		response := response.BuildErrorResponse("Failed to get all comment", http.StatusNotFound, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, getAll)
	ctx.JSON(http.StatusOK, response)

}

func (c *commentController) FindByID(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		response := response.BuildErrorResponse("Failed to parse id", http.StatusBadRequest, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	commentID := uint(id)

	photo, err := c.commentService.FindByID(commentID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to get photo", http.StatusNotFound, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, photo)
	ctx.JSON(http.StatusOK, response)
}

func (c *commentController) InsertComment(ctx *gin.Context) {

	var comment models.Comment
	contentType := ctx.Request.Header.Get("Content-Type")
	if contentType == "application/json" {
		errorBind := ctx.ShouldBindJSON(&comment)
		if errorBind != nil {
			log.Println("Error binding JSON", errorBind)
			response := response.BuildErrorResponse("Failed to bind json", http.StatusBadRequest, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	} else {
		errorBind := ctx.ShouldBind(&comment)
		if errorBind != nil {
			log.Println("Error binding JSON", errorBind)
			response := response.BuildErrorResponse("Failed to bind", http.StatusBadRequest, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["userID"].(float64))
	comment.User_id = userID

	insertComment, err := c.commentService.InsertComment(comment)
	if err != nil {
		response := response.BuildErrorResponse(err.Error(), http.StatusBadRequest, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, insertComment)
	ctx.JSON(http.StatusOK, response)
}

func (c *commentController) UpdateComment(ctx *gin.Context) {

	var comment models.Comment
	contentType := ctx.Request.Header.Get("Content-Type")
	if contentType == "application/json" {
		errorBind := ctx.ShouldBindJSON(&comment)
		if errorBind != nil {
			log.Println("Error binding JSON", errorBind)
			response := response.BuildErrorResponse("Failed to bind json", http.StatusBadRequest, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	} else {
		errorBind := ctx.ShouldBind(&comment)
		if errorBind != nil {
			log.Println("Error binding", errorBind)
			response := response.BuildErrorResponse("Failed to bind", http.StatusBadRequest, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	// get photo id & comment ID
	getCommentData := ctx.MustGet("commentData").(*models.Comment)
	comment.Photo_id = getCommentData.Photo_id
	comment.ID = getCommentData.ID

	// get user id
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["userID"].(float64))
	comment.User_id = userID

	updateComment, err := c.commentService.UpdateComment(comment)
	if err != nil {
		response := response.BuildErrorResponse(err.Error(), http.StatusBadRequest, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, updateComment)
	ctx.JSON(http.StatusOK, response)
}

func (c *commentController) DeleteComment(ctx *gin.Context) {

	getCommentData := ctx.MustGet("commentData").(*models.Comment)

	err := c.commentService.DeleteComment(getCommentData.ID)
	if err != nil {
		log.Println("Error deleting photo", err)
		response := response.BuildErrorResponse(err.Error(), http.StatusBadRequest, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK", http.StatusOK, helpers.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
