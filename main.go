package main

import (
    "github.com/gin-gonic/gin"

    "github.com/sou1991/golang_authentication_server/repository"
    "github.com/sou1991/golang_authentication_server/controller"
    "github.com/sou1991/golang_authentication_server/router"
)
var repo = repository.NewClientRepository();
var cntl = controller.NewClentController(repo);
var ro = router.NewClientRouter(cntl);

func main() {
    s := gin.Default()
    ro.CheckClent(s)
    s.Run("0.0.0.0:80")
}