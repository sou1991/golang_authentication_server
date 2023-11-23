package entity

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
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
	Code   string
	Expire time.Time
}

//In memory data
var users = []User{
	//auth@example.com/passwordを暗号化した値
	{
		Email:    "b565fcc69617dde89f5c5d82796ef890",
		Password: "a52d7fe9d86fe7255ab17ced9e084b0c",
	},
}

//In memory data
var authorized []authorizationData

func Authenticate(c *gin.Context) {
	var u User

	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	for _, v := range users {
		if v.Email != encryptUserData(u.Email) || v.Password != encryptUserData(u.Password) {
			c.JSON(http.StatusOK, gin.H{"message": "Not Found Account"})
		}else{
			t := time.Now()
			a := authorizationData{Code: "akjd783jek", Expire: t.Add(1 * time.Hour)}
	
			authorized = append(authorized, a)
	
			//https://example.com/auth?code={a.Code}
			log.Println("redirect to client server")
	
			c.JSON(http.StatusMovedPermanently, gin.H{"message": "redirect to client server"})
		}
	}
}

func encryptUserData(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	salt := "mysalt"
	io.WriteString(h, salt)

	return fmt.Sprintf("%x", h.Sum(nil))

}
