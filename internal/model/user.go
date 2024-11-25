package model

import "time"

/*
对应的建表结构
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
	is_teacher BOOLEAN NOT NULL,
    image VARCHAR(255) NOT NULL
);
*/

type User struct {
	Id        int       `gorm:"primarykey;autoIncrement"`
	Email     string    `gorm:"not null"`
	Username  string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	IsTeacher bool      `gorm:"not null"`
	Image     string    `gorm:"not null"`
}
