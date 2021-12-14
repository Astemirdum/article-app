package repository

import (
	"article/models"

	"github.com/jmoiron/sqlx"
)

const (
	userTable        = "users"
	articleTable     = "articles"
	userArticleTable = "user_article"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
}

type ArticleItem interface {
	Create(userId int, art models.Article) (int, error)
	Delete(userId, articleId int) error
	Update(userId, articleId int, upArt models.UpdateArticle) error
	SelectAll(userId int) ([]models.Article, error)
	Get(userId, articleId int) (models.Article, error)
}

type Repository struct {
	Authorization
	ArticleItem
}

func NewRepostory(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuth(db),
		ArticleItem:   NewArticlePostgres(db),
	}
}
