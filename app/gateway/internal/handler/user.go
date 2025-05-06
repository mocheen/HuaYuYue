package handler

import (
	"gateway/idl/pb/user"
	"gateway/pkg/ctl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(ctx *gin.Context) {

	var req user.LoginReq
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "绑定参数错误"))
		return
	}
	resp, err := UserClient.Login(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "UserLogin RPC服务调用错误"))
		return
	}

	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))

}

func UserRegister(ctx *gin.Context) {
	var req user.RegisterReq
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "绑定参数错误"))
		return
	}
	resp, err := UserClient.Register(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "UserRegister RPC服务调用错误"))
		return
	}

	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
}
