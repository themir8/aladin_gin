package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mirsaid-mirzohidov/aladin_gin/api"
)

func main() {
	DATABASE_URL := "postgres://mirzohidov:coder@localhost:6432/django_examples?sslmode=disable"

	db, err := sqlx.Connect("postgres", DATABASE_URL)
	if err != nil {
		log.Fatalln(err, "Database connection error")
	}
	defer db.Close()

	apiserver := api.New(db)

	apiserver.Run(":8000") // listen and serve on 0.0.0.0:8000 (for windows "localhost:8000")
}

func save(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
