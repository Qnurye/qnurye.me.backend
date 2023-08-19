package models

type Comments struct {
	CommentId  string `gorm:"column:commentId;type:SERIAL;primaryKey;" json:"commentId"`
	ParentId   string `gorm:"column:parentId;type:BIGINT UNSIGNED;" json:"parentId"`
	UserId     string `gorm:"column:userId;type:BIGINT UNSIGNED;not null;" json:"userId"`
	PostId     string `gorm:"column:postId;type:BIGINT UNSIGNED;not null;" json:"postId"`
	Content    string `gorm:"column:content;type:TEXT;not null;" json:"content"`
	UpdateTime string `gorm:"autoUpdateTime;column:updateTime;type:TIMESTAMP;" json:"updateTime"`
	CreateTime string `gorm:"autoCreateTime;column:createTime;type:TIMESTAMP;default:NOW();" json:"createTime"`
}
