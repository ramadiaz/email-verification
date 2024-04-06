package main

import (
	"email-verification/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	r := gin.Default()

	api := r.Group("/api")

	routers.CompRouters(api)

	r.Run(":" + port)
}
