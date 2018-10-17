package controller

import (
	"goim/lib/context"
	"goim/logic/db"

	"github.com/gin-gonic/gin"
)

var Engine = gin.New()

func init() {
	Engine.Use(verify)

}

func verify(c *gin.Context) {

}

func Context() *context.Context {
	return context.NewContext(db.Factoty.GetSession())
}
