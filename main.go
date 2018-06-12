package main

import (
	"./app"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	apiV1Root := router.Group("/api/v1")
	{
		titleApis := apiV1Root.Group("/titles")
		{
			titleApis.GET("", app.ReadAllTitles)
			titleApis.GET("/:id", app.ReadSingleTitle)
			titleApis.POST("", app.CreateTitle)
			titleApis.PUT("/:id", app.UpdateTitle)
			titleApis.DELETE("/:id", app.DeleteTitle)
		}

		//nameApis := apiV1Root.Group("/names")

		//searchApis := apiV1Root.Group("/search")
	}

	router.Run(":2333")

}
