package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goim/public/imerror"
)

type HandlerFunc func(*context)

func handler(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&context{Context: c})
	}
}

type context struct {
	*gin.Context
	userId int64
}

func (c *context) response(data interface{}, err error) {
	if err != nil {
		if berr, ok := err.(*imerror.BError); ok {
			c.JSON(http.StatusOK, NewWithBError(berr))
			return
		}
		c.JSON(http.StatusOK, NewWithBError(imerror.ErrUnknowError))
		return
	}
	c.Context.JSON(http.StatusOK, NewSuccess(data))
}

func (c *context) bindJson(value interface{}) error {
	err := c.ShouldBindJSON(value)
	if err != nil {
		c.JSON(http.StatusOK, NewWithBError(imerror.WrapBErrorWithData(imerror.ErrBadRequest, err)))
		return err
	}
	return nil
}
