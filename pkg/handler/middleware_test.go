package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Astemirdum/article-app/pkg/service"
	service_mocks "github.com/Astemirdum/article-app/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHandlerUserIdent(t *testing.T) {

	type mockBehavior func(r *service_mocks.MockAuthorization, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *service_mocks.MockAuthorization, token string) {
				r.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: "1",
		},
		{
			name:                 "Invalid Header Name",
			headerName:           "Auth",
			headerValue:          "Bearer token",
			token:                "token",
			mockBehavior:         func(r *service_mocks.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"empty header Authorization"}`,
		},
		{
			name:                 "Invalid Header Value",
			headerName:           "Authorization",
			headerValue:          "Bob token",
			token:                "token",
			mockBehavior:         func(r *service_mocks.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid header Authorization"}`,
		},
		{
			name:                 "Empty token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "token",
			mockBehavior:         func(r *service_mocks.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"empty header Authorization Token"}`,
		},
		{
			name:        "Parse token",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *service_mocks.MockAuthorization, token string) {
				r.EXPECT().ParseToken(token).Return(0, errors.New("invalid token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
	}
	logger := logrus.New()
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := service_mocks.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.token)

			handler := Handler{
				services: &service.Service{Authorization: auth},
				log:      logger,
			}

			r := gin.New()
			r.GET("/identity", handler.UserIdentification, func(c *gin.Context) {
				id, _ := c.Get(idKey)
				c.String(200, "%d", id)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/identity", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}

}

func TestGetUserId(t *testing.T) {
	var getContext = func(id int) *gin.Context {
		ctx := &gin.Context{}
		ctx.Set(idKey, id)
		return ctx
	}

	testTable := []struct {
		name       string
		ctx        *gin.Context
		id         int
		shouldFail bool
	}{
		{
			name:       "Empty",
			ctx:        &gin.Context{},
			shouldFail: true,
		},
		{
			ctx:  getContext(1),
			name: "Ok",
			id:   1,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			id, err := getUserId(test.ctx)
			if test.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.id, id)
		})
	}
}
