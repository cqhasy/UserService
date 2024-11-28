package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type EmailRepo interface {
	GenerateVerificationCode(ctx context.Context, expiryMinutes int)
	SendEmail(ctx context.Context, to string) error
	IsExpired(ctx context.Context) bool
}

type EmailCase struct {
	er  EmailRepo
	log *log.Helper
}

// NewUserEmailcase 创建 EmailCase 实例
func NewUserEmailcase(ur UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, log: log.NewHelper(logger)}
}
