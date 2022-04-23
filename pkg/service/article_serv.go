package service

import (
	"errors"

	"github.com/Astemirdum/article-app/models"
	"github.com/Astemirdum/article-app/pkg/repository"
)

type ArticleService struct {
	repo repository.ArticleItem
}

func NewArticleService(repo repository.ArticleItem) *ArticleService {
	return &ArticleService{
		repo: repo,
	}
}

func (a *ArticleService) Create(userId int, art models.Article) (int, error) {
	return a.repo.Create(userId, art)
}

func (a *ArticleService) Delete(userId, articleId int) error {
	return a.repo.Delete(userId, articleId)
}
func (a *ArticleService) Update(userId, articleId int, upArt models.UpdateArticle) error {
	//validation
	if upArt.Title == "" && upArt.Text == "" {
		return errors.New("empty title and text")
	}
	return a.repo.Update(userId, articleId, upArt)
}
func (a *ArticleService) SelectAll(userId int) ([]models.Article, error) {
	return a.repo.SelectAll(userId)
}

func (a *ArticleService) Get(userId, articleId int) (models.Article, error) {
	return a.repo.Get(userId, articleId)
}
