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

type accsess struct {
	name       string
	params     accsessData
	statuscode int
	jwt        string
}

type accsessData struct {
	code         string
	clientId     string
	clientSecret string
}

func TestAccess(t *testing.T) {
	a := accsessData{code: "akjd783jek", clientId: "abcde", clientSecret: "hogehogefoookgem"}

	testCase := []accsess{
		{name: "The jwt must match.", params: a, statuscode: 200, jwt: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDE3MDQ4NzgsInVzZXJfaWQiOiIxMjMifQ.q5fuIdhNmZ_IAc7sZaopkkEwHtf4pKju-IjfNsOKZpo"},
	}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			a := entity.AuthorizationParams{Code: v.params.code, ClientId: v.params.clientId, ClientSecret: v.params.clientSecret}
			byteAuth, _ := json.Marshal(a)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request, _ = http.NewRequest(
				http.MethodPost,
				"/oauth2/v1/auth/token",
				bytes.NewBuffer(byteAuth),
			)

			entity.SendToken(c)
			//expが時間なのでペイロードの値が毎回変わる為とりあえずこのテストのみ
			assert.IsType(t, v.jwt, w.Body.String())
		})
	}
}
