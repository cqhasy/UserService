package model

import "time"

// VerificationCode 验证码结构体
type VerificationCode struct {
	Code      string    // 验证码
	ExpiresAt time.Time // 过期时间
}
