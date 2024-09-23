package model

import (
    "github.com/gofrs/uuid"
    "gorm.io/gorm"
    "time"
)
type DraftStatus string

const (
	Draft     DraftStatus = "draft"
	Pending   DraftStatus = "pending"
	Published DraftStatus = "published"
)

type Post struct {
    ID       uuid.UUID   `gorm:"column:id;primaryKey"`
	PostName string      `gorm:"column:post_name"`
	UserId   string      `gorm:"column:user_id;not null;type:uuid"`
	User     User        `gorm:"foreignKey:UserId"`
	Contents  uint64      `gorm:"column:contents"`
	Draft    DraftStatus `gorm:"column:draft"`
	Likes    int64       `gorm:"column:likes"`
	Comments []string    `gorm:"column:comments;comment:Comments;type:TEXT[]"`
    CreatedAt time.Time `gorm:"column:created_at"`
    UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	post.ID = id
	return nil
}
