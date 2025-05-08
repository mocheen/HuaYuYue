package res

import (
	"gateway/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 统一响应结构体
type Response struct {
	Status  int         `json:"status"`  // 业务状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 数据
	Error   string      `json:"error"`   // 错误详情(开发调试用)
}

// 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  e.SUCCESS,
		Message: e.GetMsg(e.SUCCESS),
		Data:    data,
	})
}

// 失败响应
func Error(c *gin.Context, status int, err error) {
	c.JSON(http.StatusOK, Response{
		Status:  status,
		Message: e.GetMsg(status),
		Error:   err.Error(),
	})
	c.Abort()
}

// 带HTTP状态码的失败响应
func ErrorWithHTTPStatus(c *gin.Context, httpStatus, bizStatus int, err error) {
	c.JSON(httpStatus, Response{
		Status:  bizStatus,
		Message: e.GetMsg(bizStatus),
		Error:   err.Error(),
	})
	c.Abort()
}
