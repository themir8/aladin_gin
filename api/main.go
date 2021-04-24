package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mirsaid-mirzohidov/aladin_gin/api/v1"
)

func New() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apiV1 := router.Group("/v1")

	{
		apiV1.GET("/all", save)
		apiV1.GET("/:id", v1.GetUser)
		apiV1.POST("/save", save)
		apiV1.POST("/update", save)
		apiV1.POST("/delete", save)

	}

	return router
}

func save(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
