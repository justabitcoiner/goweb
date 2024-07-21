package db

import (
	"context"
	"fmt"
	"goweb/models"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

var conn *pgx.Conn

func Connect() {
	var err error
	conn, err = pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/goweb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}

func Disconnect() {
	conn.Close(context.Background())
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignUp(email string, password string) error {
	sql := `INSERT INTO auth_user (email, password) 
			VALUES ($1,$2)`

	hash, _ := HashPassword(password)

	_, err := conn.Exec(context.Background(), sql, email, hash)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("sign up fail")
	}

	return nil
}

func SignIn(email string, password string) (int, error) {
	sql := `SELECT id, password FROM auth_user WHERE email = $1`

	var user models.User
	err := conn.QueryRow(context.Background(), sql, email).Scan(&user.Id, &user.Password)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("email address doesn't exist")
	}

	if !CheckPasswordHash(password, user.Password) {
		return 0, fmt.Errorf("password is incorrect")
	}

	return user.Id, nil
}

// Articles
func GetArticleList() ([]models.Article, error) {
	sql := `SELECT id, user_id, title, content FROM article`

	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	article_list, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Article])
	if err != nil {
		return nil, err
	}

	return article_list, nil
}
