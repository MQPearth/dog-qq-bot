package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmptyResponse struct {
}

type ReplyResponse struct {
	Reply string `json:"reply"`
}

func EmptyOk(c *gin.Context) {
	c.JSON(http.StatusOK, EmptyResponse{})
}

func ReplyOk(c *gin.Context, reply string) {
	c.JSON(http.StatusOK, ReplyResponse{reply})
}
