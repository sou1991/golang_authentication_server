package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/entity"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testCase struct {
	name          string
	client_id     string
	response_type string
	statuscode    int
	want          string
}

func TestClient(t *testing.T) {
	badTestCases := []testCase{
		{
			name:          "missing response_type",
			client_id:     "abcde",
			response_type: "fooo",
			statuscode:    400,
			want:          `{"message":"Bad Request","reason":"invalid response type"}`,
		},
		{
			name:          "missing response_type",
			client_id:     "okgem",
			response_type: "code",
			statuscode:    400,
			want:          `{"message":"Bad Request","reason":"missing client id"}`,
		},
	}
	OkTestCase := []testCase{
		{
			name:          "succes",
			client_id:     "abcde",
			response_type: "code",
			statuscode:    200,
		},
	}

	for _, tc := range badTestCases {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, err = http.NewRequest("GET", fmt.Sprintf("/outh2/v2/auth?client_id=%v&response_type=%v", tc.client_id, tc.response_type), nil)

			if err != nil {
				t.Error(err)
			}

			entity.CheckAvailableClients(c)

			assert.Equal(t, tc.statuscode, w.Code)
			assert.Equal(t, tc.want, w.Body.String())
		})
	}

	for _, tc := range OkTestCase {
		t.Run(tc.name, func(t *testing.T) {
			var err error

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request, err = http.NewRequest("GET", fmt.Sprintf("/outh2/v2/auth?client_id=%v&response_type=%v", tc.client_id, tc.response_type), nil)

			if err != nil {
				t.Error(err)
			}

			entity.CheckAvailableClients(c)

			assert.Equal(t, tc.statuscode, w.Code)
		})
	}
}
