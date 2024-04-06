package routers

import (
	"email-verification/config"
	"email-verification/handlers"
	"email-verification/repositories"
	"email-verification/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CompRouters(api *gin.RouterGroup) {
	api.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	compRepository := repositories.NewCompRepositories(config.InitDB())
	compService := services.NewServices(compRepository)
	compHandler := handlers.NewCompHandlers(compService)

	api.POST("/get-token", compHandler.GetToken)
	api.POST("/verify-token", compHandler.VerifyToken)

}