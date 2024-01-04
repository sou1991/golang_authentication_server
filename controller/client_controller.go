package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/repository"
)

type ClientChecker interface{
	CheckClent(*gin.Context)
}

type ClientController struct{
	Cr repository.ClientChecker
}

func NewClentController(cr repository.ClientChecker) ClientChecker{
	return ClientController{Cr: cr}
}

func (cc ClientController) CheckClent(c *gin.Context){
	cc.Cr.CheckClent(c)
}
