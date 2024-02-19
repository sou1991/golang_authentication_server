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
	uuid         string
	code         string
	expire       int64
	clientId     string
	clientSecret string
}

type AuthorizationParams struct {
	Code         string `json:"code"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type response struct {
	IdToken Access `json:"access"`
}

type Access struct {
	IdToken string `json:"id_token"`
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
		uuid:         "c1a361d61cf839fe79bf6357454a88ae",
		code:         "akjd783jek",
		expire:       time.Now().Add(1 * time.Hour).Unix(),
		clientId:     "abcde",
		clientSecret: "hogehogefoookgem",
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
			a := authorizationData{code: "akjd783jek", expire: time.Now().Add(1 * time.Hour).Unix()}

			authorized = append(authorized, a)

			//https://example.com/auth?code={a.Code}
			log.Println("redirect to client server")

			c.JSON(http.StatusMovedPermanently, gin.H{"message": "redirect to client server"})
		}
	}
}

func SendIdToken(c *gin.Context) {
	var a AuthorizationParams

	if err := c.BindJSON(&a); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	for _, v := range authorized {
		if v.clientId != a.ClientId || v.clientSecret != a.ClientSecret {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "reason": "invalid client"})
			return
		} else if v.code == a.Code {
			//Payload
			claims := jwt.MapClaims{
				"user_id": v.uuid,
				"exp":     time.Now().Add(1 * time.Hour).Unix(),
			}

			//HeaderとPayloadを結合
			t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			//署名をつける
			idToken, _ := t.SignedString([]byte("my_sign"))
			r := response{IdToken: Access{idToken}}

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
