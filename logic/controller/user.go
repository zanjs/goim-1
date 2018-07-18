package controller

import (
	"github.com/gin-gonic/gin"
)

func init() {
	g := Engine.Group("/user")
	g.GET("/:id", get)
}

func get(c *gin.Context) {

}
