package handler

import (
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"role/ctl"
	"role/e"
	"role/internal/repository"
	"role/internal/repository/query"
	service "role/internal/service/pb"
	"sync"
	"time"
)

var RoleSrvIns *RoleSrv
var RoleSrvOnce sync.Once

type RoleSrv struct {
	service.UnimplementedRoleServiceServer
}

func GetRoleSrv() *RoleSrv {
	RoleSrvOnce.Do(func() {
		RoleSrvIns = &RoleSrv{}
	})
	return RoleSrvIns
}

func (s *RoleSrv) SelRole(ctx context.Context, _ *emptypb.Empty) (*service.SelRoleResp, error) {
	// 拿到当前用户id
	userInfo, success := ctl.FromContext(ctx)
	if !success {
		return nil, status.Errorf(e.ErrorUserNotExist, "用户不存在")
	}

	ur := query.UserRole
	tx := query.Q.Begin()
	// 查询用户角色
	userRole, err := tx.UserRole.Where(ur.UserId.Eq(int32(userInfo.Id))).Find()
	if err != nil {
		tx.Rollback()
		return nil, status.Errorf(e.ErrorDatabase, "数据库查询失败: %v", err)
	}
	if len(userRole) > 1 {
		tx.Rollback()
		return nil, status.Errorf(e.ErrorDatabase, "查询到多个权限: %v", err)
	}
	resp := &service.SelRoleResp{
		RoleId: userRole[0].RoleID,
	}

	err = tx.Commit()
	if err != nil {
		return nil, status.Errorf(e.ErrorDatabase, "事务提交失败: %v", err)
	}

	return resp, nil
}

func (s *RoleSrv) AddRole(ctx context.Context, req *service.AddRoleReq) (*emptypb.Empty, error) {
	em := &emptypb.Empty{}
	userRole := &repository.UserRole{
		UserId: req.UserId,
		RoleID: e.USER,
	}
	// 插入用户角色
	tx := query.Q.Begin()
	err := tx.UserRole.Create(userRole)
	if err != nil {
		tx.Rollback()
		return em, status.Errorf(e.ErrorDatabase, "数据库插入失败: %v", err)
	}
	err = tx.Commit()
	if err != nil {
		return em, status.Errorf(e.ErrorDatabase, "事务提交失败: %v", err)
	}
	return em, nil
}

func (s *RoleSrv) NewAdminAPL(ctx context.Context, req *service.NewAdminAPLReq) (*emptypb.Empty, error) {
	em := &emptypb.Empty{}
	// 检查是否为普通用户
	tx := query.Q.Begin()
	role, err := s.SelRole(ctx, nil)
	if err != nil {
		return em, err
	}
	if role.RoleId != e.USER {
		return em, status.Errorf(e.ERRORROLE, "用户已经为管理员: %v", err)
	}
	adminAPL := &repository.AdminAPL{
		APLComment: req.APLComment,
	}
	// 插入管理员申请记录
	err = tx.AdminAPL.Create(adminAPL)
	if err != nil {
		tx.Rollback()
		return em, status.Errorf(e.ErrorDatabase, "数据库插入失败: %v", err)
	}
	err = tx.Commit()
	if err != nil {
		return em, status.Errorf(e.ErrorDatabase, "事务提交失败: %v", err)
	}
	return em, nil
}

func (s *RoleSrv) SelAdminAPL(ctx context.Context, _ *emptypb.Empty) (*service.SelAdminAPLResp, error) {
	// 判断是否为超级管理员
	tx := query.Q.Begin()
	role, err := s.SelRole(ctx, nil)
	if err != nil {
		return nil, err
	}

	var aplList []*service.AdminApl
	var adminAPL []*repository.AdminAPL
	if role.RoleId == e.SUPERADMIN {
		// 查询未处理管理员申请记录
		adminAPL, err = tx.AdminAPL.Where(query.AdminAPL.Status.Eq(e.UNCONFIRMED)).Find()
		if err != nil {
			tx.Rollback()
			return nil, status.Errorf(e.ErrorDatabase, "数据库查询失败: %v", err)
		}

	} else {
		// 拿到当前用户id
		userInfo, success := ctl.FromContext(ctx)
		if !success {
			return nil, status.Errorf(e.ErrorUserNotExist, "用户不存在")
		}
		// 查询当前用户管理员申请记录
		adminAPL, err = tx.AdminAPL.Where(query.AdminAPL.UserId.Eq(uint(userInfo.Id))).Find()
	}

	for _, item := range adminAPL {
		aplList = append(aplList, &service.AdminApl{
			Id:         int32(item.ID),
			UserId:     int32(item.UserId),
			Status:     int32(item.Status),
			APLComment: item.APLComment,
			CreateAt:   timestamppb.New(item.CreatedAt),
			UpdateAt:   timestamppb.New(item.UpdatedAt),
		})
	}

	resp := &service.SelAdminAPLResp{
		AdminApl: aplList,
	}

	err = tx.Commit()
	if err != nil {
		return nil, status.Errorf(e.ErrorDatabase, "事务提交失败: %v", err)
	}
	return resp, nil
}

func (s *RoleSrv) RevAdminAPL(ctx context.Context, req *service.RevAdminAPLReq) (*emptypb.Empty, error) {
	// 判断是否为超级管理员
	tx := query.Q.Begin()
	role, err := s.SelRole(ctx, nil)
	if err != nil {
		return nil, err
	}
	if role.RoleId != e.SUPERADMIN {
		return nil, status.Errorf(e.ERRORROLE, "用户不是超级管理员: %v", err)
	}
	// 查询要更新的申请记录
	adminAPL, err := tx.AdminAPL.Where(query.AdminAPL.ID.Eq(uint(req.Id))).Find()
	if err != nil || len(adminAPL) > 1 {
		tx.Rollback()
		return nil, status.Errorf(e.ErrorDatabase, "数据库查询失败: %v", err)
	}
	// 更新管理员申请记录
	_, err = tx.AdminAPL.Where(query.AdminAPL.ID.Eq(uint(req.Id))).
		Updates(map[string]interface{}{
			"status":      req.Status,
			"rev_comment": req.REVComment,
			"updated_at":  time.Now(),
		})
	if err != nil {
		tx.Rollback()
		return nil, status.Errorf(e.ErrorDatabase, "数据库更新失败: %v", err)
	}

	// 如果审核通过，为用户添加管理员角色
	if req.Status == e.PASSED {
		_, err = tx.UserRole.Where(query.UserRole.UserId.Eq(int32(adminAPL[0].UserId))).
			Updates(map[string]interface{}{
				"role_id": e.ADMIN,
			})
		if err != nil {
			tx.Rollback()
			return nil, status.Errorf(e.ErrorDatabase, "数据库更新失败: %v", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, status.Errorf(e.ErrorDatabase, "事务提交失败: %v", err)
	}

	return nil, nil
}
