package models

type User struct {
	Id       int
	Email    string
	Password string
}

type Article struct {
	Id      int
	User    User
	UserId  int `db:"user_id"`
	Title   string
	Content string
}
