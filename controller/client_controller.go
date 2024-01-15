package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/repository"
)

type ClientChecker interface{
	CheckClent(*gin.Context)
}

type ClientController struct{
	cr repository.ClientChecker
}

func NewClentController(repo repository.ClientChecker) ClientChecker{
	return ClientController{cr: repo}
}

func (cc ClientController) CheckClent(c *gin.Context){
	cc.cr.CheckClent(c)
}
