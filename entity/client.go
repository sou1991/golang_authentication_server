package entity

import (
	"github.com/gin-gonic/gin"
	
	"net/http"
	"log"
)

// In Memory Data source
var clients = map[string]string{
	"client_id": "abcde", "client_secret": "hogehogefoookgem=",
}

type clientAuthParams struct {
	ClientId     string `form:"client_id"`
	ResponseType string `form:"response_type"`
}

func CheckAvailableClients(c *gin.Context) {
	var params clientAuthParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "reason": "missing parameter"})
		return
	}

	if params.ResponseType != "code" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "reason": "invalid response type"})
		return
	}

	if clients["client_id"] != params.ClientId {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "reason": "missing client id"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Succes"})
	}

	log.Println("return login page")
}
