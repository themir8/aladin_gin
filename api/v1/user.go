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

// func errhandler(err error, msg string) {
// 	if err != nil {
// 		log.Println(err, msg)
// 	}
// }

func DB() *sqlx.DB {
	DATABASE_URL := "postgres://mirzohidov:coder@localhost:6432/django_examples?sslmode=disable"
	db, err := sqlx.Connect("postgres", DATABASE_URL)
	if err != nil {
		log.Fatalln(err, "Database connection error")
	}
	defer db.Close()

	return db
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Fatalln(err)
		return
	}

	var user model.User

	db := DB()

	row, err := db.Queryx("SELECT id, user_id, user_name, first_name FROM aladin WHERE id = $1 LIMIT 1", i)
	if err != nil {
		log.Println(err, "Query error")
	}
	if row.Next() {
		err := row.Scan(&user.ID, &user.UserID, &user.UserName, &user.FirstName)
		if err != nil {
			log.Println(err, "Gettting user error")
		}
	}

	c.JSON(http.StatusOK, user)
}
