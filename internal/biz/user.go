package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type User struct {
	Id        int
	Email     string
	Username  string
	Password  string
	IsTeacher bool
}

type UserLogin struct {
	Email    string
	Username string
}

// UserRepo 定义用户相关的数据库操作接口
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindById(ctx context.Context, id int) (*User, error)
}

type UserUsecase struct {
	ur  UserRepo
	log *log.Helper
}

// NewUserUsecase 创建 UserUsecase 实例
func NewUserUsecase(ur UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (*UserLogin, error) {
	// 从仓库中根据邮箱获取用户信息
	user, err := uc.ur.FindByEmail(ctx, email)
	if err != nil {
		// 如果查询不到用户，返回错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// 校验密码是否正确
	if user.Password != password {
		return nil, ErrInvalidPassword
	}

	// 返回登录成功信息
	return &UserLogin{
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (uc *UserUsecase) Register(ctx context.Context, username string, email string, password string) (*User, error) {
	// 校验用户是否已存在
	existingUser, err := uc.ur.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	// 创建用户
	user := &User{
		Email:     email,
		Username:  username,
		Password:  password, // 注意：密码还未实现加密
		IsTeacher: false,
	}
	if err := uc.ur.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
