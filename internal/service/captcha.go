package service

import (
	"UserService/internal/model"
	"crypto/tls"
	"fmt"
	"gopkg.in/yaml.v3"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"time"
)

// GenerateVerificationCode 生成一个6位的数字验证码
func GenerateVerificationCode(expiryMinutes int) *model.VerificationCode {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000 // 生成一个100000到999999之间的随机数
	return &model.VerificationCode{
		Code:      strconv.Itoa(code),
		ExpiresAt: time.Now().Add(time.Duration(expiryMinutes) * time.Minute), // 设置过期时间
	}
}

type EmailConfig struct {
	Email struct {
		From       string `yaml:"from"`
		Password   string `yaml:"password"`
		SMTPServer string `yaml:"smtp_server"`
		SMTPHost   string `yaml:"smtp_host"`
	} `yaml:"email"`
}

// NewEmailConfigs NewMysqlConfigs 读取并解析 YAML 配置文件
func NewEmailConfigs(configFilePath string) (*EmailConfig, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return &EmailConfig{}, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var configs EmailConfig
	// 解析 YAML 文件
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&configs)
	if err != nil {
		return &EmailConfig{}, err
	}

	return &configs, nil
}

// SendEmail 发送邮件函数
func SendEmail(config *EmailConfig, to string) error {
	from := config.Email.From
	password := config.Email.Password // 邮箱授权码
	smtpServer := config.Email.SMTPServer

	// 设置 PlainAuth
	auth := smtp.PlainAuth("", from, password, config.Email.SMTPHost)

	// 创建 tls 配置
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.qq.com",
	}

	// 连接到 SMTP 服务器
	conn, err := tls.Dial("tcp", smtpServer, tlsconfig)
	if err != nil {
		return fmt.Errorf("TLS 连接失败: %v", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, "smtp.qq.com")
	if err != nil {
		return fmt.Errorf("SMTP 客户端创建失败: %v", err)
	}
	defer client.Quit()

	// 使用 auth 进行认证
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("认证失败: %v", err)
	}

	// 设置发件人和收件人
	if err = client.Mail(from); err != nil {
		return fmt.Errorf("发件人设置失败: %v", err)
	}
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("收件人设置失败: %v", err)
	}

	// 写入邮件内容
	wc, err := client.Data()
	if err != nil {
		return fmt.Errorf("数据写入失败: %v", err)
	}
	defer wc.Close()

	subject := "CCNU-EDU-LLM"
	code := GenerateVerificationCode(1)
	body := `
		<h1>Verification Code</h1>
		<p>Your verification code is: <strong>` + code.Code + `</strong></p>
		<p>This verification code is valid for 15 minutes</p>
		<p>If you are not doing it yourself, please ignore it !</p>
	`
	msg := []byte("From: Sender Name <" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body)

	_, err = wc.Write(msg)
	if err != nil {
		return fmt.Errorf("消息发送失败: %v", err)
	}

	return nil
}

// IsExpired 检查验证码是否过期
func IsExpired(v *model.VerificationCode) bool {
	return time.Now().After(v.ExpiresAt)
}
