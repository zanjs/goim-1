package controller

import (
	"goim/logic/service"

	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

const OK = 200

func init() {
	g := Engine.Group("/user")

	g.GET("/:id", get)
}

func get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
	}
	user, err := service.NewUserService().Get(id)
	c.JSON(OK, user)
}
