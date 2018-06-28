package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func APIRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to TripleZ/IMDB-Backend API root !",
		"note": "For more API documentation, please refer to https://github.com/Triple-Z/IMDB-Backend#api-references .",
	})
}