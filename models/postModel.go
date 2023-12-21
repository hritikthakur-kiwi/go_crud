package model

import (
	"gorm.io/gorm"
)

type DraftStatus string

const (
	Draft     DraftStatus = "draft"
	Pending   DraftStatus = "pending"
	Published DraftStatus = "published"
)

type Post struct {
	gorm.Model
	PostName string      `gorm:"column:post_name"`
	UserId   uint        `gorm:"column:user_id;not null"` // Foreign key referencing User.ID
	User     User        `gorm:"foreignKey:UserId"`
	Content  uint64      `gorm:"column:content"`
	Draft    DraftStatus `gorm:"column:draft"`
	Likes    int64       `gorm:"column:likes"`
	Comments []string    `gorm:"column:comments;comment:Comments;type:TEXT[]"`
}
