package model

type User struct {
	ID        int    `json:id`
	UserID    uint   `json:user_id`
	UserName  string `json:username`
	FirstName string `json:first_name`
}
