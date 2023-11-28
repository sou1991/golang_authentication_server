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
}

func (authRepository) Auth(c ...*gin.Context){
	entity.Authenticate(c[0])
}