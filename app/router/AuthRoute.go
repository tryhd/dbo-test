package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tryhd/dbo-test/app/controllers"
	"github.com/tryhd/dbo-test/app/models"
)

func LoginRoute(route *gin.Engine) {
	m := models.NewAuthModels()
	m.Init()
	c := controllers.NewAuthController(m)
	authRoutes := route.Group("api/v1/auth")
	{
		authRoutes.POST("/login", c.Login)
		// authRoutes.Use(middlewares.JwtAuthMiddleware())
		authRoutes.POST("/register", c.Register)
	}
}
