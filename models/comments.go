package models

import (
	"errors"
	"qnurye/qnurye.me/pkg/db"
	"time"
)

type Comment struct {
	CommentId  uint      `gorm:"column:commentId;type:SERIAL;primaryKey;" json:"commentId"`
	ParentId   uint      `gorm:"column:parentId;type:BIGINT UNSIGNED;" json:"parentId"`
	UserId     uint      `gorm:"column:userId;type:BIGINT UNSIGNED;not null;" json:"userId"`
	PostId     uint      `gorm:"column:postId;type:BIGINT UNSIGNED;not null;" json:"postId"`
	Content    string    `gorm:"column:content;type:TEXT;not null;" json:"content"`
	UpdateTime time.Time `gorm:"autoUpdateTime;column:updateTime;type:TIMESTAMP;" json:"updateTime"`
	CreateTime time.Time `gorm:"autoCreateTime;column:createTime;type:TIMESTAMP;default:NOW();" json:"createTime"`
}

var CommentNotFound = errors.New("comment not found")

func (c Comment) Create(parentId uint, userId uint, postId uint, content string) (*Comment, error) {
	d := db.Get()

	parent, err := new(Comment).GetById(parentId)
	if err != nil {
		return nil, err
	}
	user, err := new(User).GetById(userId)
	if err != nil {
		return nil, err
	}
	post, err := new(Post).GetById(postId)
	if err != nil {
		return nil, err
	}

	c.ParentId = parent.CommentId
	c.UserId = user.UserId
	c.PostId = post.PostId
	c.Content = content

	if err := d.Create(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func (c Comment) GetById(i uint) (*Comment, error) {
	err := db.Get().First(&c, i)
	if err != nil {
		return &c, CommentNotFound
	}

	return &c, nil
}
