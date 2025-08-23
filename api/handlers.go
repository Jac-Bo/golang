package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArticleHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusAccepted, gin.H{
		"message": GetArticles(c),
	})
}

func LikeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Article Like Handler",
	})
}
