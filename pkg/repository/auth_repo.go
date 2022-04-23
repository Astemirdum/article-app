package repository

import (
	"fmt"
	"github.com/Astemirdum/article-app/models"

	"github.com/jmoiron/sqlx"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db}
}

func (a *Auth) CreateUser(usr models.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, hash_password, email) VALUES ($1, $2, $3) RETURNING id", userTable)
	row := a.db.QueryRow(query, &usr.Name, &usr.Password, &usr.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *Auth) GetUser(email, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 AND hash_password=$2", userTable)
	err := a.db.Get(&user, query, email, password)
	return user, err
}
