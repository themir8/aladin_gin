package model

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID        int64  `json:id`
	UserID    int    `json:user_id`
	UserName  string `json:user_name`
	FirstName string `json:first_name`
}

func (p *User) SaveUser(db *sqlx.DB) {
	err := db.QueryRowx("INSERT INTO aladin(user_id, user_name, first_name) values($1, $2, $3) RETURNING id", p.UserID, p.UserName, p.FirstName).Scan(&p.ID)
	if err != nil {
		log.Println(err, " Bla Bla2")
		return
	}
	fmt.Println("person saved")
}

func (p *User) GetUser(db *sqlx.DB, id int64) (*User, error) {
	row, err := db.Queryx("SELECT id, user_id, user_name, first_name FROM aladin WHERE id = $1 LIMIT 1", id)
	if err != nil {
		log.Println(err, " Bla Bla2")
	}
	if row.Next() {
		err := row.Scan(&p.ID, &p.UserID, &p.UserName, &p.FirstName)
		if err != nil {
			log.Println(err, " Bla Bla2")
		}
	}

	return p, err
}

func (p *User) DeleteUser(db *sqlx.DB) {
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM aladin WHERE id = $1", p.ID)
	tx.Commit()
}

func ListUser(db *sqlx.DB) (persons []User) {
	rows, _ := db.Queryx("SELECT * FROM aladin")
	for rows.Next() {
		var person User
		err := rows.Scan(&person.ID, &person.UserID, &person.UserName, &person.FirstName)
		if err != nil {
			log.Println(err, " Bla Bla2")
		}
		persons = append(persons, person)
	}
	fmt.Println(persons)
	fmt.Println("person list")
	return
}
