package controller

import (
	"github.com/gin-gonic/gin"
)

var Engine = gin.New()

func init() {
	Engine.Use(verify)

}

func verify(c *gin.Context) {

}
