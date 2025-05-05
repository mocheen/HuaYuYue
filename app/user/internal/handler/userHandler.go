package handler

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"user/internal/repository"
	"user/internal/repository/query"

	service "user/internal/service/pb"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
	service.UnimplementedUserServiceServer
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (s *UserSrv) Register(ctx context.Context, req *service.RegisterReq) (*service.RegisterResp, error) {
	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "密码加密失败: %v", err)
	}

	u := query.User
	tx := query.Q.Begin()

	// 检查邮箱是否已经存在
	count, err := tx.User.Where(u.Email.Eq(req.Email)).Count()
	if err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "数据库查询失败: %v", err)
	}
	if count > 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.AlreadyExists, "该邮箱已注册")
	}

	// 创建用户对象
	user := &repository.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		UserName: req.Username,
	}

	// 插入用户到数据库
	err = tx.User.Create(user)
	if err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "用户创建失败: %v", err)
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "事务提交失败: %v", err)
	}

	// 返回响应
	resp := &service.RegisterResp{
		UserId: int32(user.ID),
	}
	return resp, nil
}

func (s *UserSrv) Login(ctx context.Context, req *service.LoginReq) (*service.LoginResp, error) {
	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "密码加密失败: %v", err)
	}

	u := repository.User{
		UserName: "1",
		Password: "1",
		Email:    "1",
	}
	repository.DB.Create(&u)

	// 示例返回 Token
	return &service.LoginResp{
		Token: string(hashedPassword),
	}, nil
}
