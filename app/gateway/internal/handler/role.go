package handler

import (
	"errors"
	"gateway/idl/role"
	"gateway/pkg/e"
	"gateway/pkg/res"
	"github.com/gin-gonic/gin"
)

func SelRole(ctx *gin.Context) {
	resp, err := RoleClient.SelRole(ctx.Request.Context(), nil)
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

	_, err := RoleClient.AddRole(ctx.Request.Context(), &req)
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

	_, err := RoleClient.NewAdminAPL(ctx.Request.Context(), &req)
	if err != nil {
		res.Error(ctx, e.ServiceError, err)
		return
	}
	res.Success(ctx, "success")
}

func SelAdminAPL(ctx *gin.Context) {
	resp, err := RoleClient.SelAdminAPL(ctx.Request.Context(), nil)
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

	_, err := RoleClient.RevAdminAPL(ctx.Request.Context(), &req)
	if err != nil {
		res.Error(ctx, e.ServiceError, err)
		return
	}
	res.Success(ctx, "success")
}
