package main

import (
	"fmt"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Unable to load .env file")
	}

	r := gin.Default()

	// db := database.InitialDB()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET(fmt.Sprintf("/%s/%s/ping", os.Getenv("API_PREFIX"), os.Getenv("API_VERSION")), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "service started",
			"status":  "Service running ✅ Test update",
		})
	})

	v1 := r.Group("/" + os.Getenv("API_PREFIX") + "/" + os.Getenv("API_VERSION"))

	v1.Static("storage", "./assets")

	r.Run(":" + os.Getenv("APP_PORT"))
}
