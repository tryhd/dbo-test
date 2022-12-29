package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tryhd/dbo-test/app/config/middleware"
	"github.com/tryhd/dbo-test/app/controllers"
	"github.com/tryhd/dbo-test/app/models"
)

func OrderRoute(route *gin.Engine) {
	m := models.NewOrderModels()
	m.Init()
	c := controllers.NewOrderController(m)
	authRoutes := route.Group("api/v1/order")
	{
		authRoutes.Use(middleware.JwtAuthMiddleware())
		authRoutes.POST("/create", c.RegisterOrder)
		authRoutes.GET("/all", c.GetAllOrder)
		authRoutes.GET("/:id", c.DetailOrder)
		authRoutes.GET("/", c.FindOrder)
		authRoutes.PUT("/:id", c.UpdateOrder)
		authRoutes.DELETE("/:id", c.DeleteOrder)

	}
}
