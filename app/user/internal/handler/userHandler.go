package handler

import (
	"context"
	"google.golang.org/grpc/status"
	"sync"
	"user/e"
	"user/internal/repository"
	"user/internal/repository/query"
	"user/jwt"
	"user/util"

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
	hashedPassword := util.HashWithSalt(req.Password, req.Username)

	u := query.User
	tx := query.Q.Begin()

	// 检查邮箱是否已经存在
	count, err := tx.User.Where(u.Email.Eq(req.Email)).Count()
	if err != nil {
		tx.Rollback()
		return nil, status.Errorf(e.ErrorDatabase, "数据库查询失败: %v", err)
	}
	if count > 0 {
		tx.Rollback()
		return nil, status.Errorf(e.ErrorUserAlreadyExist, "该邮箱已注册")
	}

	// 检查用户名是否已经存在
	count, err = tx.User.Where(u.UserName.Eq(req.Username)).Count()
	if err != nil {
		tx.Rollback()
		return nil, status.Errorf(e.ErrorDatabase, "数据库查询失败: %v", err)
	}
	if count > 0 {
		tx.Rollback()
		return nil, status.Errorf(e.ErrorUserAlreadyExist, "该用户名已存在")
	}

	// 创建用户对象
	user := &repository.User{
		Email:    req.Email,
		Password: hashedPassword,
		UserName: req.Username,
	}

	// 插入用户到数据库
	err = tx.User.Create(user)
	if err != nil {
		tx.Rollback()
		return nil, status.Errorf(e.ErrorDatabase, "用户创建失败: %v", err)
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		return nil, status.Errorf(e.ErrorDatabase, "事务提交失败: %v", err)
	}

	// 返回响应
	resp := &service.RegisterResp{
		UserId: int32(user.ID),
	}
	return resp, nil
}

func (s *UserSrv) Login(ctx context.Context, req *service.LoginReq) (*service.LoginResp, error) {
	// 密码加密
	hashedPassword := util.HashWithSalt(req.Password, req.Username)

	u := query.User
	tx := query.Q.Begin()
	user, err := tx.User.Where(u.UserName.Eq(req.Username), u.Password.Eq(hashedPassword)).First()
	if user == nil {
		user, err = tx.User.Where(u.Email.Eq(req.Username)).First()
		if user != nil {
			hashedPassword = util.HashWithSalt(req.Password, user.UserName)
			user, err = tx.User.Where(u.Email.Eq(req.Username), u.Password.Eq(hashedPassword)).First()
		}
	}
	if err != nil {
		tx.Rollback()
		return nil, status.Errorf(e.ErrorDatabase, "用户名或密码错误: %v", err)
	}

	token, err := jwt.GenerateToken(int64(user.ID))
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, status.Errorf(e.ErrorDatabase, "事务提交失败: %v", err)
	}

	return &service.LoginResp{
		Token: token,
	}, nil
}
