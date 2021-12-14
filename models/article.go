package models

import "time"

type Article struct {
	Id      int       `json:"id" db:"id"`
	Title   string    `json:"title" db:"title" binding:"required"`
	Text    string    `json:"text" db:"thesis"`
	PubTime time.Time `json:"-" db:"pub_time"`
}

type UserArticle struct {
	Id        int `db:"id"`
	UserId    int `db:"user_id"`
	ArticleId int `db:"article_id"`
}

type UpdateArticle struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
