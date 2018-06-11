package main

import (
	"github.com/gin-gonic/gin"
	"./app"
)

func main() {

	router := gin.Default()

	apiV1Root := router.Group("/api/v1")
	{
		titleApis := apiV1Root.Group("/titles")
		{
			titleApis.GET("", app.ReadAllTitles)
		}

		//nameApis := apiV1Root.Group("/names")

		//searchApis := apiV1Root.Group("/search")
	}
}
