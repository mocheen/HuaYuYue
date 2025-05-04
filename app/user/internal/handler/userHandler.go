package handler

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"user/internal/repository"
	"user/internal/repository/query"
	service "user/internal/service/pb"
)

type UserService struct {
	service.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Register(ctx context.Context, req *service.RegisterReq) (*service.RegisterResp, error) {
	// 实现注册逻辑
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := query.User
	count, err := query.Q.User.Where(u.Email.Eq(req.Email)).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("email already exists")
	}

	user := &repository.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		UserName: req.Username,
	}

	tx := query.Q.Begin()
	err = query.Q.User.Create(user)
	if err != nil {
		return nil, err
	}
	resp := &service.RegisterResp{
		UserId: int32(user.ID),
	}

	if err != nil {
		tx.Rollback()
	}

	tx.Commit()
	return resp, nil

}

func (s *UserService) Login(ctx context.Context, req *service.LoginReq) (*service.LoginResp, error) {
	// 实现登录逻辑
	fmt.Printf("Received Login Request: %+v\n", req)

	// 示例返回 Token
	return &service.LoginResp{
		Token: "abc123xyz",
	}, nil
}
