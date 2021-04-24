package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mirsaid-mirzohidov/aladin_gin/model"
)

func GetUser(c *gin.Context, db *sqlx.DB) {
	id := c.Param("id")

	_, err := uuid.Parse(id)

	if err != nil {
		log.Fatalln(err)
		return
	}

	var user model.User

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Fatalln(err)
		return
	}

	response, err := user.GetUser(db, i)
	// log.Println(response)

	// if err != nil {
	// 	log.Fatalln(err, "Getting Permission by id")
	// 	return
	// }

	c.JSON(http.StatusOK, response)
}
