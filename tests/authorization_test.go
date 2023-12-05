package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/entity"
	"github.com/stretchr/testify/assert"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type authentication struct {
	name       string
	params     user
	statusCode int
}

type user struct {
	email    string
	password string
}

func TestAuth(t *testing.T) {
	authTests := []authentication{
		{
			name:       "Authorization success",
			params:     user{email: "auth@example.com", password: "password"},
			statusCode: http.StatusMovedPermanently,
		},
	}

	for _, v := range authTests {
		t.Run(v.name, func(t *testing.T) {
			u := entity.authorizationData{Email: v.params.email, Password: v.params.password}
			byteUser, _ := json.Marshal(u)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request, _ = http.NewRequest(
				http.MethodPost,
				"/oauth2/v1/auth/login",
				bytes.NewBuffer(byteUser),
			)

			entity.Authenticate(c)
			assert.Equal(t, http.StatusMovedPermanently, w.Code)
		})
	}
}
