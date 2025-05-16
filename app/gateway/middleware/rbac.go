package middleware

import (
	"errors"
	"gateway/internal/handler"
	"gateway/pkg/e"
	"gateway/pkg/res"
	"github.com/gin-gonic/gin"
)

func RBAC() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, err := handler.RoleClient.SelRole(c.Request.Context(), nil)
		if err != nil {
			res.Error(c, e.ServiceError, err)
			c.Abort()
			return
		}
		if role.RoleId == e.USER {
			res.Error(c, e.ErrorRole, errors.New("权限不足"))
		}

		c.Next()
	}
}
