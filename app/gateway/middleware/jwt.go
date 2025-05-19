package middleware

import (
	"errors"
	"gateway/pkg/ctl"
	"gateway/pkg/e"
	"gateway/pkg/jwt"
	"gateway/pkg/res"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			res.ErrorWithHTTPStatus(c, 401, e.ErrorAuthTokenInvalid,
				errors.New("authorization header is missing or invalid"))
			return
		}
		token = token[7:]
		claims, err := jwt.ParseToken(token)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		}
		if code != e.SUCCESS {
			res.ErrorWithHTTPStatus(c, 401, code, err)
			return
		}
		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.UserID}))
		ctl.InitUserInfo(c.Request.Context())
		c.Next()
	}
}
