package main

import (
    "github.com/gin-gonic/gin"
    "github.com/sou1991/golang_authentication_server/repository"
    "github.com/sou1991/golang_authentication_server/controller"
    "github.com/sou1991/golang_authentication_server/router"
)
var cr = repository.NewClientRepository()
var cc = controller.NewClentController(cr)
var cro = router.NewClientRouter(cc)

var ar = repository.NewAuthRepository()
var ac = controller.NewAuthController(ar)
var aro = router.NewAuthRouter(ac)

func main() {
    s := gin.Default()
    //Client Cheak
    //cro.CheckClent(s)
    //Auth
    //aro.Auth(s)
    s.Run("0.0.0.0:80")
}