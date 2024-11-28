package service

import (
	v1 "UserService/api/userapi/v1"
	"UserService/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type EmailService struct {
	v1.UnimplementedEmailServer

	uc  *biz.EmailUsecase
	log *log.Helper
}

func NewEmailService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *EmailService) SendVerificationCode(ctx context.Context, req *v1.SendVerificationCodeRequest) (reply *v1.SendVerificationCodeReply, err error) {
	rv, err := s.uc.Login(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.SendVerificationCodeReply{
		Message: "",
	}, nil
}
