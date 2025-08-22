package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// serve allows the game to
	router.Use(static.Serve("/", static.LocalFile(
		"./views/js", true,
	)))

	api := router.Group("/api") //group of routes under the var api and route "/api"
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api.GET("/articles")
	api.POST("/articles/like/:artID")

	router.Run(":3000")
}
