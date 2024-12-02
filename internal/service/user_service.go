package service

import (
	v1 "UserService/api/userapi/v1"
	"UserService/internal/biz"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.UserReply, err error) {
	rv, err := s.uc.Login(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username:  rv.Username,
			Email:     rv.Email,
			IsTeacher: rv.IsTeacher,
		},
		Token: rv.Token,
	}, nil
}

func (s *UserService) Register(ctx context.Context, req *v1.RegisterRequest) (reply *v1.UserReply, err error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password, req.User.ConfirmPassword, req.User.IsTeacher, req.VerificationCode)
	if err != nil {
		return nil, err
	}

	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username:  u.Username,
			Email:     u.Email,
			IsTeacher: u.IsTeacher,
		},
		Token: u.Token,
	}, nil
}

func (s *UserService) SendVerificationCode(ctx context.Context, req *v1.SendVerificationCodeRequest) (reply *v1.SendVerificationCodeReply, err error) {
	// 发送邮件
	err = s.uc.SendEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to send email: %v", err)
	}

	return &v1.SendVerificationCodeReply{Message: "Verification code sent successfully"}, nil
}
