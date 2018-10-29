package controller

import (
	"goim/logic/db"
	"goim/public/ctx"

	"github.com/gin-gonic/gin"
)

var Engine = gin.New()

func init() {
	Engine.Use(verify)

}

// verify 权限校验
func verify(c *gin.Context) {
	//token := c.GetHeader("token")

}

func Context() *ctx.Context {
	return ctx.NewContext(db.Factoty.GetSession())
}
