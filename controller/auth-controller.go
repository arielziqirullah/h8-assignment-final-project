package controller

import (
	"h8-assignment-final-project/helpers"
	"h8-assignment-final-project/models"
	"h8-assignment-final-project/service"
	"h8-assignment-final-project/web/request"
	"h8-assignment-final-project/web/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	userService service.UserService
	jwtService  service.JwtService
}

func NewAuthController(userService service.UserService, jwtService service.JwtService) AuthController {
	return &authController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {

	var loginRequest request.LoginRequest
	contentType := ctx.Request.Header.Get("Content-Type")
	if contentType != "application/json" {
		errorBind := ctx.ShouldBindJSON(&loginRequest)
		if errorBind != nil {
			log.Println("Error binding JSON", errorBind)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Error binding JSON",
			})
			return
		}
	} else {
		errorBind := ctx.ShouldBind(&loginRequest)
		if errorBind != nil {
			log.Println("Error binding JSON", errorBind)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Error binding",
			})
			return
		}
	}

	verifyCredential, err := c.userService.VerifyCredential(loginRequest.Username, loginRequest.Password)
	if err != nil {
		log.Println("Error verifying credential", err)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  false,
			"message": "Error verifying credential",
		})
		return
	}

	if value, ok := verifyCredential.(models.User); ok {
		generateToken, expToken, err := c.jwtService.GenerateToken(value.ID, value.Email)
		if err != nil {
			response := response.BuildErrorResponse("Failed to generate token", http.StatusUnprocessableEntity, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
			return
		}

		loginBuild := response.LoginResponseDTO{
			AccessToken: generateToken,
			ExpiresIn:   expToken,
			TokenType:   "Bearer",
		}

		response := response.BuildResponse(true, "OK", http.StatusOK, loginBuild)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := response.BuildErrorResponse("Please check your credential", http.StatusUnprocessableEntity, helpers.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
}

func (c *authController) Register(ctx *gin.Context) {

	var user models.User
	contentType := ctx.Request.Header.Get("Content-Type")
	if contentType != "application/json" {
		errorBind := ctx.ShouldBindJSON(&user)
		if errorBind != nil {
			log.Println("Error binding JSON", errorBind)
			response := response.BuildErrorResponse("Failed to bind json", http.StatusBadRequest, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	} else {
		errorBind := ctx.ShouldBind(&user)
		if errorBind != nil {
			log.Println("Error binding", errorBind)
			response := response.BuildErrorResponse("Failed to bind", http.StatusBadRequest, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	if !c.userService.MinAge(user.Age) {
		log.Println("Age must be greater than 8")
		response := response.BuildErrorResponse("Age must be greater than 8", http.StatusBadRequest, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	isDuplicate, err := c.userService.IsDuplicateEmail(user.Email)
	if err != nil {
		log.Println("Error checking duplicate email", err)
		response := response.BuildErrorResponse("Error checking duplicate email", http.StatusUnprocessableEntity, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	if isDuplicate {
		response := response.BuildErrorResponse("Email already exists", http.StatusConflict, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	} else {
		createdUser, errInsert := c.userService.InsertUser(&user)
		if errInsert != nil {
			log.Println("Error inserting user", errInsert)
			response := response.BuildErrorResponse(errInsert.Error(), http.StatusUnprocessableEntity, helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
			return
		} else {
			ctx.JSON(http.StatusCreated, gin.H{
				"status":  true,
				"message": "User created successfully",
				"data":    createdUser,
			})
		}
	}

}
