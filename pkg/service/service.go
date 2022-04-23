package service

import (
	"github.com/Astemirdum/article-app/models"
	"github.com/Astemirdum/article-app/pkg/repository"
)

//go:generate go run github.com/golang/mock/mockgen -source=service.go -destination=mocks/mock.go

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
