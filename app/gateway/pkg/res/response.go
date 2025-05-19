package res

import (
	"fmt"
	"gateway/pkg/e"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
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
	var jsonData []byte
	var err error

	// 检查是否是 protobuf 消息
	if msg, ok := data.(proto.Message); ok {
		// 使用 protobuf 的 JSON 序列化器
		marshaler := protojson.MarshalOptions{
			EmitUnpopulated: true, // 确保零值字段也被序列化
		}
		jsonData, err = marshaler.Marshal(msg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 手动构建响应
		c.Data(http.StatusOK, "application/json", []byte(fmt.Sprintf(
			`{"status":%d,"message":"%s","data":%s}`,
			e.SUCCESS,
			e.GetMsg(e.SUCCESS),
			jsonData,
		)))
		return
	}

	// 非 protobuf 消息走原来的逻辑
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
