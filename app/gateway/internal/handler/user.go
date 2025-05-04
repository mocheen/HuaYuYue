package handler

import (
	"context"
	"gateway/idl/pb/user"

	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	var req user.RegisterReq
	PanicIfUserError(ctx.Bind(&req))

	userService := ctx.Keys["user"].(user.UserServiceClient)
	resp, err := userService.Register(context.Background(), &req)
	PanicIfUserError(err)

	ctx.JSON(http.StatusOK, resp)
}

func UserLogin(ctx *gin.Context) {
	var req user.LoginReq
	PanicIfUserError(ctx.Bind(&req))

	userService := ctx.Keys["user"].(user.UserServiceClient)
	resp, err := userService.Login(context.Background(), &req)
	PanicIfUserError(err)

	ctx.JSON(http.StatusOK, resp)
}
