package handler

import (
	"bytes"
	"errors"
	"github.com/Astemirdum/article-app/pkg/service"
	service_mocks "github.com/Astemirdum/article-app/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Astemirdum/article-app/models"
)

func TestHandler_SingUp(t *testing.T) {
	type input struct {
		body string
		user models.User
	}
	type response struct {
		expectedCode int
		expectedBody string
	}
	type mockBehavior func(r *service_mocks.MockAuthorization, user models.User)

	tests := []struct {
		name         string
		mockBehavior mockBehavior
		input        input
		response     response
	}{
		{
			name: "ok",
			mockBehavior: func(r *service_mocks.MockAuthorization, user models.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			input: input{
				body: `{"Name": "lol", "password": "kek", "email": "lol@kek"}`,
				user: models.User{
					Name:     "lol",
					Password: "kek",
					Email:    "lol@kek",
				},
			},
			response: response{
				expectedCode: http.StatusOK,
				expectedBody: `{"userId":1}`,
			},
		},
		{
			name:         "ko. invalid input",
			mockBehavior: func(r *service_mocks.MockAuthorization, user models.User) {},
			input: input{
				body: `{"Name": "lol", "email": "lol@kek"}`,
				user: models.User{
					Name:  "lol",
					Email: "lol@kek",
				},
			},
			response: response{
				expectedCode: http.StatusBadRequest,
				expectedBody: `{"message":"singUp invalid input"}`,
			},
		},
		{
			name: "ko. internal error",
			mockBehavior: func(r *service_mocks.MockAuthorization, user models.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("db query fail"))
			},
			input: input{
				body: `{"Name": "", "password": "kek", "email": "lol@kek"}`,
				user: models.User{
					Name:     "",
					Password: "kek",
					Email:    "lol@kek",
				},
			},
			response: response{
				expectedCode: http.StatusInternalServerError,
				expectedBody: `{"message":"singUp db query fail"}`,
			},
		},
	}
	log := logrus.New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repo := service_mocks.NewMockAuthorization(c)
			tt.mockBehavior(repo, tt.input.user)

			h := &Handler{
				services: &service.Service{
					Authorization: repo,
				},
				log: log,
			}
			g := gin.New()
			g.POST("/auth/sign-up", h.SingUp)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/auth/sign-up",
				bytes.NewBufferString(tt.input.body))

			g.ServeHTTP(w, r)

			require.Equal(t, tt.response.expectedCode, w.Code)
			require.Equal(t, tt.response.expectedBody, w.Body.String())
		})
	}
}
