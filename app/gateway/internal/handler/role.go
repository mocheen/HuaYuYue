package handler

import (
	"errors"
	"fmt"
	"gateway/pkg/ctl"
	"google.golang.org/grpc/metadata"

	"gateway/idl/role"

	"gateway/pkg/e"
	"gateway/pkg/res"
	"github.com/gin-gonic/gin"
)

func SelRole(ctx *gin.Context) {
	userInfo, err := ctl.GetUserInfo(ctx.Request.Context())
	grpcCtx := metadata.NewOutgoingContext(
		ctx.Request.Context(),
		metadata.Pairs("x-user-id", fmt.Sprintf("%d", userInfo.Id)),
	)
	resp, err := RoleClient.SelRole(grpcCtx, nil)
	if err != nil {
		res.Error(ctx, e.ServiceError, err)
		return
	}
	res.Success(ctx, resp)
}

func AddRole(ctx *gin.Context) {
	var req role.AddRoleReq
	if err := ctx.Bind(&req); err != nil {
		res.Error(ctx, e.ErrorInvalidParams, errors.New("绑定参数错误"))
		return
	}

	_, err := RoleClient.AddRole(ctx, &req)
	if err != nil {
		res.Error(ctx, e.ServiceError, err)
		return
	}
	res.Success(ctx, "success")
}

func NewAdminAPL(ctx *gin.Context) {
	var req role.NewAdminAPLReq
	if err := ctx.Bind(&req); err != nil {
		res.Error(ctx, e.ErrorInvalidParams, errors.New("绑定参数错误"))
		return
	}
	userInfo, err := ctl.GetUserInfo(ctx.Request.Context())
	grpcCtx := metadata.NewOutgoingContext(
		ctx.Request.Context(),
		metadata.Pairs("x-user-id", fmt.Sprintf("%d", userInfo.Id)),
	)

	_, err = RoleClient.NewAdminAPL(grpcCtx, &req)
	if err != nil {
		res.Error(ctx, e.ServiceError, err)
		return
	}
	res.Success(ctx, "success")
}

func SelAdminAPL(ctx *gin.Context) {
	userInfo, err := ctl.GetUserInfo(ctx.Request.Context())
	grpcCtx := metadata.NewOutgoingContext(
		ctx.Request.Context(),
		metadata.Pairs("x-user-id", fmt.Sprintf("%d", userInfo.Id)),
	)
	resp, err := RoleClient.SelAdminAPL(grpcCtx, nil)
	if err != nil {
		res.Error(ctx, e.ServiceError, err)
		return
	}
	res.Success(ctx, resp)
}

func RevAdminAPL(ctx *gin.Context) {
	var req role.RevAdminAPLReq
	if err := ctx.Bind(&req); err != nil {
		res.Error(ctx, e.ErrorInvalidParams, errors.New("绑定参数错误"))
		return
	}
	userInfo, err := ctl.GetUserInfo(ctx.Request.Context())
	grpcCtx := metadata.NewOutgoingContext(
		ctx.Request.Context(),
		metadata.Pairs("x-user-id", fmt.Sprintf("%d", userInfo.Id)),
	)

	_, err = RoleClient.RevAdminAPL(grpcCtx, &req)
	if err != nil {
		res.Error(ctx, e.ServiceError, err)
		return
	}
	res.Success(ctx, "success")
}
