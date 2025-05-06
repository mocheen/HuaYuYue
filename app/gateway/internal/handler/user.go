package handler

import (
	"errors"
	"gateway/idl/pb/user"
	"gateway/pkg/e"
	"gateway/pkg/res"
	"github.com/gin-gonic/gin"
)

func UserLogin(ctx *gin.Context) {

	var req user.LoginReq
	if err := ctx.Bind(&req); err != nil {
		res.Error(ctx, e.ErrorInvalidParams, errors.New("绑定参数错误"))
		return
	}
	resp, err := UserClient.Login(ctx, &req)

	if err != nil {
		res.Error(ctx, e.ServiceError, err)
		return
	}

	res.Success(ctx, resp)

}

func UserRegister(ctx *gin.Context) {
	var req user.RegisterReq
	if err := ctx.Bind(&req); err != nil {
		res.Error(ctx, e.ErrorInvalidParams, errors.New("绑定参数错误"))
		return
	}
	resp, err := UserClient.Register(ctx, &req)
	if err != nil {
		res.Error(ctx, e.ServiceError, err)
		return
	}

	res.Success(ctx, resp)
}
