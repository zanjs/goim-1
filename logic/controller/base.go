package controller

import (
	"goim/logic/db"
	"goim/public/ctx"

	"goim/logic/service"
	"goim/public/imerror"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

var Engine = gin.New()

func init() {
	Engine.Use(handler(verify))

}

const (
	keyDeviceId = "device_id"
	keyUserId   = "user_id"
)

// verify 权限校验
func verify(c *context) {
	gin.Logger()
	deviceIdStr := c.GetHeader("device_id")
	token := c.GetHeader("token")

	if c.HandlerName() == "/device" {
		return
	}

	deviceId, err := strconv.ParseInt(deviceIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, NewWithBError(imerror.ErrUnauthorized))
		c.Abort()
		return
	}
	userId, err := service.AuthService.Auth(Context(), deviceId, token)
	if err != nil {
		c.JSON(http.StatusOK, NewWithBError(imerror.ErrUnauthorized))
		c.Abort()
		return
	}
	c.Keys = make(map[string]interface{}, 2)
	c.Keys[keyDeviceId] = deviceId
	if c.HandlerName() != "/user" && c.HandlerName() != "/user/signin" {
		if userId == 0 {
			c.JSON(http.StatusOK, NewWithBError(imerror.ErrDeviceNotBindUser))
			c.Abort()
			return
		}
		c.Keys[keyUserId] = userId
	}
	c.Next()
}

func Context() *ctx.Context {
	return ctx.NewContext(db.Factoty.GetSession())
}
