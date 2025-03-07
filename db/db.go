package db

import (
	"context"
	"fmt"
	"goweb/models"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var pool *pgxpool.Pool

func Connect() {
	var err error
	pool, err = pgxpool.New(context.Background(), "postgres://postgres:123456@localhost:5432/goweb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}

func Disconnect() {
	pool.Close()
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

	_, err := pool.Exec(context.Background(), sql, email, hash)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("sign up fail")
	}

	return nil
}

func SignIn(email string, password string) (int, error) {
	sql := `SELECT id, password FROM auth_user WHERE email = $1`

	var user models.User
	err := pool.QueryRow(context.Background(), sql, email).Scan(&user.Id, &user.Password)
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

	rows, err := pool.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	article_list, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Article])
	if err != nil {
		return nil, err
	}

	return article_list, nil
}

func GetArticleDetail(id int) (*models.Article, error) {
	sql := `SELECT id, user_id, title, content FROM article WHERE id = $1`

	rows, err := pool.Query(context.Background(), sql, id)
	if err != nil {
		return nil, err
	}

	article_list, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Article])
	if err != nil {
		return nil, err
	}

	if len(article_list) == 0 {
		return nil, nil
	}

	return &article_list[0], nil
}

func CreateNewArticle(userId int, title string, content string) error {
	sql := `INSERT INTO article (user_id, title, content) VALUES ($1, $2, $3)`

	_, err := pool.Exec(context.Background(), sql, userId, title, content)
	if err != nil {
		return err
	}
	return nil
}

func UpdateArticle(userId int, articleId int, title string, content string) error {
	sql := `UPDATE article SET title = $1, content = $2 WHERE id = $3 AND user_id = $4`

	_, err := pool.Exec(context.Background(), sql, title, content, articleId, userId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(userId int, articleId int) error {
	sql := `DELETE FROM article WHERE id = $1 AND user_id = $2`

	_, err := pool.Exec(context.Background(), sql, articleId, userId)
	if err != nil {
		return err
	}
	return nil
}
