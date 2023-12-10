package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/entity"
)

func NewAuthRepository() Authenticator {
	return authRepository{}
}

type authRepository struct{}

type Authenticator interface{
	Auth(...*gin.Context)
	Access(...*gin.Context)
}

func (authRepository) Auth(c ...*gin.Context){
	entity.Authenticate(c[0])
}

func (authRepository) Access(c ...*gin.Context){
	entity.SendIdToken(c[0])
}