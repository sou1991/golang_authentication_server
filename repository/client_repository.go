package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/entity"
)

type ClientRepository struct{}

func NewClientRepository() ClientChecker{
	return ClientRepository{}
}

type ClientChecker interface{
	CheckClent(*gin.Context)
}

func (ClientRepository) CheckClent(c *gin.Context){
	entity.CheckAvailableClients(c)
}
