package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Req struct {
	Code int64
	Msg  string
	Data interface{}
}

func Re(c *gin.Context, code int64, msg string, data interface{}) {
	httpCode := http.StatusOK
	if code > 20000 {
		httpCode = http.StatusBadGateway
	}
	c.JSON(httpCode, &Req{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	return
}
