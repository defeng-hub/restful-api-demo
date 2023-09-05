package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

const (
	ERROR   = -1
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context, err ...error) {
	resp := Response{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	if len(err) > 0 {
		resp.Error = err[0].Error()
	}
	// 开始时间
	c.JSON(http.StatusOK, resp)
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context, err error) {
	Result(ERROR, map[string]interface{}{}, "operation failed", c, err)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
