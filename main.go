package main

import (
	routes "golang-jwt-project/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	//from internal package routes
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context){
		c.JSON(200,gin.H{"success":"Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context){
		c.JSON(200,gin.H{"success":"Access granted for api-2"})
	})

	router.RUN(":" + port)
}
