package controller

import (
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func RegisterPingController(engine *gin.RouterGroup) {
	var group = engine.Group("/v1")
	group.GET("/ping", Ping)
}
