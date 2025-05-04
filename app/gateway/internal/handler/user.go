package handler

import (
	"context"
	service "gateway/internal/service/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	var req service.RegisterReq
	PanicIfUserError(ctx.Bind(&req))

	userService := ctx.Keys["user"].(service.UserServiceClient)
	resp, err := userService.Register(context.Background(), &req)
	PanicIfUserError(err)

	ctx.JSON(http.StatusOK, resp)
}

func UserLogin(ctx *gin.Context) {
	var req service.LoginReq
	PanicIfUserError(ctx.Bind(&req))

	userService := ctx.Keys["user"].(service.UserServiceClient)
	resp, err := userService.Login(context.Background(), &req)
	PanicIfUserError(err)

	ctx.JSON(http.StatusOK, resp)
}
