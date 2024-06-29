package utils

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// Options 请求
func Options() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,"+
			"X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization,Token")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func ByteToStruct(jsonStr []byte, data interface{}) {
	decoder := json.NewDecoder(bytes.NewReader(jsonStr))
	if err := decoder.Decode(data); err != nil {
		panic(err)
	}
}
