package models

import (
	"errors"
	"qnurye/qnurye.me/pkg/db"
	"time"
)

type User struct {
	UserId       uint      `gorm:"column:userId;type:SERIAL;primaryKey;" json:"userId"`
	Nickname     string    `gorm:"column:nickname;type:VARCHAR(255);not null;" json:"nickname"`
	Email        string    `gorm:"column:email;type:VARCHAR(255);not null;unique" json:"email"`
	Password     string    `gorm:"column:password;type:VARCHAR(255);not null;" json:"password"`
	Website      string    `gorm:"column:website;type:VARCHAR(255);not null;" json:"website"`
	Subscribed   bool      `gorm:"column:subscribed;type:BOOLEAN;not null;default:false;" json:"subscribed"`
	UpdateTime   time.Time `gorm:"autoUpdateTime;column:updateTime;type:TIMESTAMP;" json:"updateTime"`
	RegisterTime time.Time `gorm:"autoCreateTime;column:registerTime;type:TIMESTAMP;default:now();" json:"registerTime"`
}

func (u User) Create(nickname string, email string, pwd string, website string) (*User, error) {
	u.Nickname = nickname
	u.Email = email

	// TODO: 完成用户通过邮箱密码创建部分

	d := db.Get()

	if err := d.Create(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

var UserNotFound = errors.New("comment not found")

func (u User) GetById(i uint) (*User, error) {
	err := db.Get().First(&u, i)
	if err != nil {
		return &u, UserNotFound
	}

	return &u, nil
}
