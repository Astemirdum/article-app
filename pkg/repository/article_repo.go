package repository

import (
	"fmt"
	"strings"

	"github.com/Astemirdum/article-app/models"
	"github.com/jmoiron/sqlx"
)

type ArticlePostgres struct {
	db *sqlx.DB
}

func NewArticlePostgres(db *sqlx.DB) *ArticlePostgres {
	return &ArticlePostgres{db: db}
}

func (a *ArticlePostgres) Create(userId int, art models.Article) (int, error) {
	var id int

	tx, err := a.db.Begin()
	if err != nil {
		return 0, nil
	}
	query := fmt.Sprintf("INSERT INTO %s (title, thesis) VALUES ($1, $2) RETURNING id", articleTable)
	row := tx.QueryRow(query, art.Title, art.Text)

	if err = row.Scan(&id); err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	query = fmt.Sprintf("INSERT INTO %s (user_id, article_id) VALUES ($1, $2)", userArticleTable)
	_, err = tx.Exec(query, userId, id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (a *ArticlePostgres) Delete(userId, articleId int) error {

	query := fmt.Sprintf("DELETE FROM %s a USING %s ua WHERE a.id=ua.article_id AND a.id=$1 AND ua.user_id=$2 ", articleTable, userArticleTable)
	_, err := a.db.Exec(query, articleId, userId)
	return err
}
func (a *ArticlePostgres) Update(userId, articleId int, upArt models.UpdateArticle) error {

	setArgs := make([]string, 0)
	args := make([]any, 0)
	c := 1
	if upArt.Title != "" {
		setArgs = append(setArgs, fmt.Sprintf("title=$%d", c))
		args = append(args, upArt.Title)
		c++
	}
	if upArt.Text != "" {
		setArgs = append(setArgs, fmt.Sprintf("thesis=$%d", c))
		args = append(args, upArt.Text)
		c++
	}

	setStr := strings.Join(setArgs, ", ")
	query := fmt.Sprintf("UPDATE %s a SET %s FROM %s ua  WHERE a.id=ua.article_id AND a.id=$%d AND ua.user_id=$%d",
		articleTable, setStr, userArticleTable, c, c+1)
	args = append(args, articleId, userId)

	_, err := a.db.Exec(query, args...)
	return err
}
func (a *ArticlePostgres) SelectAll(userId int) ([]models.Article, error) {
	var articles []models.Article
	query := fmt.Sprintf(`
	SELECT a.id, a.title, a.thesis, a.pub_time FROM %s a 
	INNER JOIN %s ua 
	ON a.id = ua.article_id
	WHERE ua.user_id = $1`,
		articleTable, userArticleTable)

	err := a.db.Select(&articles, query, userId)
	return articles, err
}

func (a *ArticlePostgres) Get(userId, articleId int) (models.Article, error) {
	var article models.Article
	query := fmt.Sprintf(`
	SELECT a.id, a.title, a.thesis, a.pub_time FROM %s a 
	INNER JOIN %s ua
	ON a.id=ua.article_id
	WHERE ua.user_id = $1 AND a.id = $2
	`, articleTable, userArticleTable)

	err := a.db.Get(&article, query, userId, articleId)

	return article, err
}
