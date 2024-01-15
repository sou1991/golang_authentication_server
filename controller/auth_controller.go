package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/repository"
)

func NewAuthController(repo repository.Authenticator) Authenticator{
	return authController{ra: repo}
}

type authController struct{
	ra repository.Authenticator
}

type Authenticator interface{
	Auth(*gin.Context)
	Access(*gin.Context)
}

func(ac authController) Auth(c *gin.Context){
	ac.ra.Auth(c)
}

func(ac authController) Access(c *gin.Context){
	ac.ra.Access(c)
}
