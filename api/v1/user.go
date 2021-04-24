package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mirsaid-mirzohidov/aladin_gin/model"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Fatalln(err)
		return
	}
	var user model.User

	response := model.UserInter.GetUser(i)

	c.JSON(http.StatusOK, user)
}

func SaveUser(c *gin.Context) {
	var err error
	DATABASE_URL := "postgres://mirzohidov:coder@localhost:6432/django_examples?sslmode=disable"
	db, err = sqlx.Connect("postgres", DATABASE_URL)

	if err != nil {
		log.Fatalln(err, "Database connection error")
	}

	defer db.Close()

	var user model.User

	err = db.QueryRowx("INSERT INTO aladin(user_id, user_name, first_name) values($1, $2, $3) RETURNING id", user.UserID, user.UserName, user.FirstName).Scan(&user.ID)
	if err != nil {
		log.Println(err, "User saving error")
		return
	}

	c.JSON(http.StatusOK, user)
}
