package routers

import (
	"h8-assignment-final-project/controller"
	"h8-assignment-final-project/database"
	"h8-assignment-final-project/middlewares"
	"h8-assignment-final-project/repository"
	"h8-assignment-final-project/service"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                    *gorm.DB                         = database.SetUpDatabaseConnection()
	userRepository        repository.UserRepository        = repository.NewUserRepository(db)
	photoRepository       repository.PhotoRepository       = repository.NewPhotoRepository(db)
	commentRepository     repository.CommentRepository     = repository.NewCommentRepository(db)
	socialMediaRepository repository.SocialMediaRepository = repository.NewSocialMediaRepository(db)
	photoService          service.PhotoService             = service.NewPhotoService(photoRepository)
	commentService        service.CommentService           = service.NewCommentService(commentRepository)
	socalMediaService     service.SocialMediaService       = service.NewSocialMediaService(socialMediaRepository)
	userService           service.UserService              = service.NewUserService(userRepository)
	jwtService            service.JwtService               = service.NewJwtService()
	authController        controller.AuthController        = controller.NewAuthController(userService, jwtService)
	photoController       controller.PhotoController       = controller.NewPhotoController(photoService)
	commentController     controller.CommentController     = controller.NewCommentController(commentService)
	socialMediaController controller.SocialMediaController = controller.NewSocialMediaController(socalMediaService)

	middleWareAuthentication middlewares.AuthenticationMiddleware = middlewares.NewAuthenticationMiddleware(jwtService)
	middleWareAuthorization  middlewares.AuthorizationMiddleware  = middlewares.NewAuthorizationMiddleware(db)
)

func Run() *gin.Engine {
	defer database.CloseDatabaseConnection(db)

	routes := gin.Default()

	routes.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":         "OK",
			"message":        "Hacktive 8 Final Project",
			"golang_version": runtime.Version(),
		})
	})

	// Register route
	routesV1 := routes.Group("v1")
	{
		authRoutes := routesV1.Group("/api/auth")
		{
			authRoutes.POST("/login", authController.Login)
			authRoutes.POST("/register", authController.Register)
		}

		photoRoutes := routesV1.Group("/api/photo")
		{
			photoRoutes.Use(middleWareAuthentication.Authentication())
			photoRoutes.GET("/", photoController.FindAll)
			photoRoutes.GET("/:id", photoController.FindByID)
			photoRoutes.POST("/", photoController.InsertPhoto)

			// Authorization Photo Route
			photoRoutes.Use(middleWareAuthorization.PhotoAuthorization())
			photoRoutes.PUT("/:id", photoController.UpdatePhoto)
			photoRoutes.DELETE("/:id", photoController.DeletePhoto)
		}

		commentRoutes := routesV1.Group("/api/comment")
		{
			commentRoutes.Use(middleWareAuthentication.Authentication())
			commentRoutes.GET("/", commentController.FindAll)
			commentRoutes.GET("/:id", commentController.FindByID)
			commentRoutes.POST("/", commentController.InsertComment)

			// Authorization Comment Route
			commentRoutes.Use(middleWareAuthorization.CommentAuthorization())
			commentRoutes.PUT("/:id", commentController.UpdateComment)
			commentRoutes.DELETE("/:id", commentController.DeleteComment)
		}

		socialMediaRoutes := routesV1.Group("/api/social-media")
		{
			socialMediaRoutes.Use(middleWareAuthentication.Authentication())
			socialMediaRoutes.GET("/", socialMediaController.FindAll)
			socialMediaRoutes.GET("/:id", socialMediaController.FindByID)
			socialMediaRoutes.POST("/", socialMediaController.InsertSocialMedia)

			// Authorization Social Media Route
			socialMediaRoutes.Use(middleWareAuthorization.SocialMediaAuthorization())
			socialMediaRoutes.PUT("/:id", socialMediaController.UpdateSocialMedia)
			socialMediaRoutes.DELETE("/:id", socialMediaController.DeleteSocialMedia)
		}
	}

	if os.Getenv("RUNNING_PORT") == "" {
		log.Println("listening and serving on default port :8080")
		routes.Run()
		return routes
	} else {
		log.Println("listening and serving on port " + os.Getenv("RUNNING_PORT"))
		routes.Run(":" + os.Getenv("RUNNING_PORT"))
		return routes
	}
}
