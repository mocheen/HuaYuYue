package service

import (
	"app/user/biz/dal/model"
	"app/user/biz/dal/query"
	user "app/user/kitex_gen/user"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
}

// NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	exist, err := query.Q.User.Where(query.User.Email.Eq(req.Email)).First()
	if err != nil {
		// 查询出错
		return nil, err
	}
	if exist != nil {
		// 如果存在该 email 用户，返回错误
		return nil, errors.New("该邮箱已被注册")
	}

	// 创建新用户
	newUser := &model.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Username: req.Username,
	}

	err = query.Q.User.Create(newUser)
	if err != nil {
		return
	}
	resp = &user.RegisterResp{
		UserId: int32(newUser.ID),
	}

	return
}
