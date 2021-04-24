package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	v1 "github.com/mirsaid-mirzohidov/aladin_gin/api/v1"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apiV1 := router.Group("/v1")

	{
		apiV1.GET("/all", save)
		apiV1.GET("/:id", v1.GetUser(db))
		apiV1.POST("/save", save)
		apiV1.POST("/update", save)
		apiV1.POST("/delete", save)

	}

	return router
}
