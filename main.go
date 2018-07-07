package IMDB_Backend

import (
	"./app"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	if app.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	if app.AllowCORS {
		//router.Use(cors.Default())  // There has a bug in this cors version
		router.Use(CORS())
	}

	apiV1Root := router.Group("/api/v1")
	{
		apiV1Root.GET("", app.APIRoot)

		titleApis := apiV1Root.Group("/titles")
		{
			titleApis.GET("", app.ReadAllTitles)
			titleApis.GET("/:id", app.ReadSingleTitle)
			//titleApis.POST("", app.CreateTitle)
			//titleApis.PUT("/:id", app.UpdateTitle)
			//titleApis.DELETE("/:id", app.DeleteTitle)

			titleApis.GET("/:id/details", app.ReadTitleDetails)
		}

		nameApis := apiV1Root.Group("/names")
		{
			nameApis.GET("", app.ReadAllNames)
			nameApis.GET("/:id", app.ReadSingleName)
			//nameApis.POST("", app.CreateName)
			//nameApis.PUT("/:id", app.UpdateName)
			//nameApis.DELETE("/:id", app.DeleteName)
		}

		fuzzySearchApis := apiV1Root.Group("/search")
		{
			fuzzySearchApis.GET("/titles", app.FuzzySearchTitles)
			fuzzySearchApis.GET("/names", app.FuzzySearchNames)
		}

		advancedSearchApis := apiV1Root.Group("/advanced_search")
		{
			advancedSearchApis.GET("/titles", app.AccurateSearchTitles)
			advancedSearchApis.GET("/names", app.AccurateSearchNames)
		}

		principalApis := apiV1Root.Group("/principals")
		{
			principalApis.GET("/:id", app.ReadSinglePrincipal)
			//principalApis.POST("", app.CreatePrincipal)
			//principalApis.PUT("/:id", app.UpdatePrincipal)
			//principalApis.DELETE("/:id", app.DeletePrincipal)
		}
	}

	router.Run(":" + app.PortNumber)

}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", app.CORSAllowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
