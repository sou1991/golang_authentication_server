package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/router"
)

func main() {
	s := gin.Default()
	router.Router(s)
	s.Run("0.0.0.0:80")
}
