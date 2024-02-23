package http

import (
	"github.com/dmytrodemianchuk/go-auth/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoint(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/sign-up", h.SignUp)
		authEndpoints.POST("/sign-in", h.SignIn)
	}
}
