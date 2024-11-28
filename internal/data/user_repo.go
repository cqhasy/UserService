package data

import (
	"UserService/internal/biz"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

/*
// SQL 创建语句
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_teacher BOOLEAN NOT NULL DEFAULT FALSE
);
*/

// User 是数据库表的模型
type User struct {
	Id        int    `gorm:"primarykey;autoIncrement;column:id"`
	Email     string `gorm:"type:varchar(255);not null;uniqueIndex;column:email"`
	Username  string `gorm:"type:varchar(255);not null;column:username"`
	Password  string `gorm:"type:varchar(255);not null;column:password"`
	IsTeacher bool   `gorm:"not null;default:false;column:is_teacher"`
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo 创建 userRepo 实例
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u userRepo) FindByEmail(ctx context.Context, email string) (*biz.User, error) {
	dbUser := &User{}
	if err := u.data.db.WithContext(ctx).Where("email = ?", email).First(dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &biz.User{
		Id:        dbUser.Id,
		Email:     dbUser.Email,
		Username:  dbUser.Username,
		Password:  dbUser.Password,
		IsTeacher: dbUser.IsTeacher,
	}, nil
}

func (u userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	dbUser := &User{}
	if err := u.data.db.WithContext(ctx).Where("username = ?", username).First(dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &biz.User{
		Id:        dbUser.Id,
		Email:     dbUser.Email,
		Username:  dbUser.Username,
		Password:  dbUser.Password,
		IsTeacher: dbUser.IsTeacher,
	}, nil
}

func (u userRepo) FindById(ctx context.Context, id int) (*biz.User, error) {
	dbUser := &User{}
	if err := u.data.db.WithContext(ctx).First(dbUser, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &biz.User{
		Id:        dbUser.Id,
		Email:     dbUser.Email,
		Username:  dbUser.Username,
		Password:  dbUser.Password,
		IsTeacher: dbUser.IsTeacher,
	}, nil
}

func (u userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	dbUser := &User{
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password, // 实际中需加密
		IsTeacher: user.IsTeacher,
	}
	return u.data.db.WithContext(ctx).Create(dbUser).Error
}
