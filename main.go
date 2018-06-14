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

			titleApis.GET("/:id/detail", app.ReadTitleDetails)
		}

		nameApis := apiV1Root.Group("/names")
		{
			nameApis.GET("", app.ReadAllNames)
			nameApis.GET("/:id", app.ReadSingleName)
			nameApis.POST("", app.CreateName)
			nameApis.PUT("/:id", app.UpdateName)
			nameApis.DELETE("/:id", app.DeleteName)
		}

		searchApis := apiV1Root.Group("/search")
		{
			searchApis.GET("/titles", app.FuzzySearchTitles)
			searchApis.POST("/titles", app.AccurateSearchTitles)
			searchApis.GET("/names", app.FuzzySearchNames)
			searchApis.POST("/names", app.AccurateSearchNames)
		}
	}

	router.Run(":2333")

}
