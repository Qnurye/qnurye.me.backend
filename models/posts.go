package models

import (
	"errors"
	"qnurye/qnurye.me/pkg/db"
	"time"
)

type Post struct {
	PostId     uint      `gorm:"column:postId;type:SERIAL;primaryKey;" json:"postId"`
	Title      string    `gorm:"column:title;type:VARCHAR(151);not null;" json:"title"`
	Summary    string    `gorm:"column:summary;type:VARCHAR(1801);" json:"summary"`
	Content    string    `gorm:"column:content;type:MEDIUMTEXT;not null;" json:"content"`
	WordCount  int       `gorm:"column:wordCount;type:BIGINT UNSIGNED;not null;default:0;" json:"wordCount"`
	ViewCount  int       `gorm:"column:viewCount;type:BIGINT UNSIGNED;not null;default:0;" json:"viewCount"`
	LikeCount  int       `gorm:"column:likeCount;type:BIGINT UNSIGNED;not null;default:0;" json:"likeCount"`
	UpdateTime time.Time `gorm:"autoUpdateTime;column:updateTime;type:TIMESTAMP;" json:"updateTime"`
	UploadTime time.Time `gorm:"autoCreateTime;column:uploadTime;type:TIMESTAMP;default:now();" json:"uploadTime"`
}

func (p Post) Create(title string, sum string, content string) (*Post, error) {
	d := db.Get()

	p.Title = title
	p.Summary = sum
	p.Content = content

	if err := d.Create(&p).Error; err != nil {
		return nil, err
	}

	return &p, nil
}

var PostNotFound = errors.New("comment not found")

func (p Post) GetById(i uint) (*Post, error) {
	err := db.Get().First(&p, i)
	if err != nil {
		return &p, PostNotFound
	}

	return &p, nil
}
