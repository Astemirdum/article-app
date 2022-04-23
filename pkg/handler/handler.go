package handler

import (
	"net/http"

	"github.com/Astemirdum/article-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{
		services: srv,
	}
}

type Index struct {
	EndPoint    string `json:"endpoint"`
	Description string `json:"description"`
}

func (h *Handler) NewRouter() *gin.Engine {
	router := gin.New()

	router.GET("/", IndexHandler)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SingUp)
		auth.POST("/sign-in", h.SingIn)
	}
	api := router.Group("/api", h.UserIdentification) // identification
	{
		article := api.Group("/article")
		{
			article.GET("/", h.GetAllArticles)      // get all
			article.GET("/:id", h.GetOneArticle)    // get one
			article.POST("/", h.CreateArticle)      // add one
			article.PUT("/:id", h.UpdateArticle)    // update one
			article.DELETE("/:id", h.DeleteArticle) // delete one
		}
	}

	return router
}

func IndexHandler(c *gin.Context) {
	index := []Index{
		{
			EndPoint:    "/sign-up",
			Description: "POST 'email' 'password' 'name'?",
		},
		{
			EndPoint:    "/sign-in",
			Description: "POST 'email' 'password'",
		},
		{
			EndPoint:    "/api/article/",
			Description: "GET all",
		},
		{
			EndPoint:    "/api/article/",
			Description: "POST add article  'text' 'title'",
		},
		{
			EndPoint:    "/api/article/:id",
			Description: "PUT update 'text'? 'title'?",
		},
		{
			EndPoint:    "/api/article/:id",
			Description: "GET one",
		},
		{
			EndPoint:    "/api/article/:id",
			Description: "DELETE one",
		},
	}
	c.JSON(http.StatusOK,
		gin.H{"index Endpoints in json format": index},
	)
}
