package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Astemirdum/article-app/models"
	"github.com/Astemirdum/article-app/pkg/repository"
	"github.com/golang-jwt/jwt"
)

const (
	salt     = "kjvbbe8392dsn"
	signKey  = "oiqwc#b891FEWFWSD"
	tokenTTL = 30 * time.Minute
)

type mytokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = genPasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func (a *AuthService) ParseToken(accessToken string) (int, error) {
	jwtToken, err := jwt.ParseWithClaims(accessToken, &mytokenClaims{},
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("wrong signing method")
			}
			return []byte(signKey), nil
		})
	if err != nil {
		return 0, err
	}
	myclaims, ok := jwtToken.Claims.(*mytokenClaims)
	if !ok {
		return 0, errors.New("token claims -> wrong suit for *mytokenClaims")
	}
	return myclaims.UserId, nil
}

func (a *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := a.repo.GetUser(email, genPasswordHash(password))
	if err != nil {
		return "", err
	}
	claims := mytokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).UTC().Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
		},
		UserId: user.Id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	return token.SignedString([]byte(signKey))
}

func genPasswordHash(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
