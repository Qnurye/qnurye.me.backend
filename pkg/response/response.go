package response

import "github.com/gin-gonic/gin"

type response struct {
	Code code   `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Success(c *gin.Context, data ...any) *gin.Context {
	c.JSON(200, response{
		Code: SUCCESS,
		Data: data,
		Msg:  "",
	})
	return c
}

func Failure(c *gin.Context, code code, msg string) *gin.Context {
	c.JSON(200, response{
		Code: code,
		Data: nil,
		Msg:  msg,
	})
	return c
}
