package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Router() {
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load enf file!")
	}
	r := gin.New()
	LoginRoute(r)
	CustomerRoute(r)
	OrderRoute(r)
	// r.Use(configs.CORSMiddleware())

	r.Run(":" + os.Getenv("APP_PORT"))

}
