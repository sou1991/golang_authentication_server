package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authorizationData struct {
	Uuid				string
	Code         string
	Expire       int64
	ClientId     string
	ClientSecret string
}

type AuthorizationParams struct {
	Code         string `json:"code"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type Response struct {
	AccessToken Access `json:"access"`
}

type Access struct {
	Token string `json:"access_token"`
}

// In memory data
var users = []User{
	//auth@example.com/passwordを暗号化した値
	{
		Email:    "b565fcc69617dde89f5c5d82796ef890",
		Password: "a624f8389727f067e9b89933e4f77d1f",
	},
}

// In memory data
var authorized = []authorizationData{
	{
		Uuid:		"c1a361d61cf839fe79bf6357454a88ae",
		Code:         "akjd783jek",
		Expire:       time.Now().Add(1 * time.Hour).Unix(),
		ClientId:     "abcde",
		ClientSecret: "hogehogefoookgem",
	},
}

func Authenticate(c *gin.Context) {
	var u User

	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	for _, v := range users {
		if v.Email != encryptUserData(u.Email) || v.Password != encryptUserData(u.Password) {
			c.JSON(http.StatusOK, gin.H{"message": "Not Found Account"})
		} else {
			a := authorizationData{Code: "akjd783jek", Expire: time.Now().Add(1 * time.Hour).Unix()}

			authorized = append(authorized, a)

			//https://example.com/auth?code={a.Code}
			log.Println("redirect to client server")

			c.JSON(http.StatusMovedPermanently, gin.H{"message": "redirect to client server"})
		}
	}
}

func SendToken(c *gin.Context) {
	var a AuthorizationParams

	if err := c.BindJSON(&a); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	for _, v := range authorized {
		if v.ClientId != a.ClientId || v.ClientSecret != a.ClientSecret {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "reason": "invalid client"})
			return
		} else if v.Code == a.Code {
			//Payload
			claims := jwt.MapClaims{
				"user_id": v.Uuid,
				"exp":     time.Now().Add(1 * time.Hour).Unix(),
			}

			//HeaderとPayloadを結合
			t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			//署名をつける
			accessToken, _ := t.SignedString([]byte("my_sign"))
			r := Response{AccessToken: Access{accessToken}}

			c.JSON(http.StatusOK, r)
			return
		} else {
			continue
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
}

func encryptUserData(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	salt := "mysalt"
	io.WriteString(h, salt)

	return fmt.Sprintf("%x", h.Sum(nil))
}
