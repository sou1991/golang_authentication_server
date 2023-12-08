package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/repository"
)

func NewAuthController(ra repository.Authenticator) Authenticator{
	return authController{Ra: ra}
}

type authController struct{
	Ra repository.Authenticator
}

type Authenticator interface{
	Auth(...*gin.Context)
	Access(...*gin.Context)
}

func(ac authController) Auth(c ...*gin.Context){
	ac.Ra.Auth(c[0])
}

func(ac authController) Access(c ...*gin.Context){
	ac.Ra.Access(c[0])
}