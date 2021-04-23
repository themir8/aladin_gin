package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	v1 := r.Group("/v1/user")
	{
		v1.GET("/ping", pong)
		// v1.POST("/submit", submitEndpoint)
		// v1.POST("/read", readEndpoint)
	}

	r.Run(":8000") // listen and serve on 0.0.0.0:8000 (for windows "localhost:8000")
}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
