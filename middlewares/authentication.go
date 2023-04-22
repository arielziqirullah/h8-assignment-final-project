package middlewares

import (
	"h8-assignment-final-project/helpers"
	"h8-assignment-final-project/service"
	"h8-assignment-final-project/web/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationMiddleware interface {
	Authentication() gin.HandlerFunc
}

type authenticationMiddleware struct {
	jwtService service.JwtService
}

func NewAuthenticationMiddleware(jwtService service.JwtService) AuthenticationMiddleware {
	return &authenticationMiddleware{
		jwtService: jwtService,
	}
}

func (s *authenticationMiddleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		verifyToken, err := s.jwtService.ValidateToken(header)
		_ = verifyToken

		if err != nil {
			log.Println("Error validating token", err)
			response := response.BuildErrorResponse("Unautheticated", http.StatusUnauthorized, helpers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if verifyToken != nil {
			c.Set("userData", verifyToken)
			c.Next()
		} else {
			log.Println("User Data Not Found")
			response := response.BuildErrorResponse("User Data Not Found", http.StatusUnauthorized, helpers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
