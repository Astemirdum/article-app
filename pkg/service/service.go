package service

import (
	"article/models"
	"article/pkg/repository"
)

type Service struct {
	Authorization
	ArticleItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		ArticleItem:   NewArticleService(repo.ArticleItem),
	}
}

type Authorization interface {
	CreateUser(models.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ArticleItem interface {
	Create(userId int, art models.Article) (int, error)
	Delete(userId, articleId int) error
	Update(userId, articleId int, upArt models.UpdateArticle) error
	SelectAll(userId int) ([]models.Article, error)
	Get(userId, articleId int) (models.Article, error)
}
