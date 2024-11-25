package service

import (
	"fmt"
	"testing"
	"time"
)

func TestSendEmail(t *testing.T) {
	var email = "2967522781@qq.com"
	var config, err = NewEmailConfigs("../../configs/email.yaml")
	if err != nil {
		t.Errorf("获取邮箱配置失败，error:%v", err)
	}
	// 发送验证码邮件
	err = SendEmail(config, email)
	if err != nil {
		t.Errorf("发送邮件失败:%v", err)
	}
}

func TestIsExpired(t *testing.T) {
	// 生成一个有效期为1分钟的验证码
	var code = GenerateVerificationCode(1)

	// 每秒输出提示，持续2秒
	fmt.Print("Testing before expiration: ")
	for i := 1; i <= 2; i++ {
		fmt.Printf("\rTesting before expiration: %ds elapsed", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Print("\r") // 清理提示

	// 检查验证码是否过期
	if IsExpired(code) != false {
		t.Errorf("Verification.IsExpired()=%v, want %v", true, false)
	}

	// 每秒输出提示，持续65秒，模拟等待到过期
	fmt.Print("Testing after expiration: ")
	for i := 1; i <= 65; i++ {
		fmt.Printf("\rTesting after expiration: %ds elapsed", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Print("\r") // 清理提示

	// 再次检查验证码是否过期
	if IsExpired(code) != true {
		t.Errorf("After expiration Verification.IsExpired()=%v, want %v", false, true)
	}
}
