package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"column:id;primaryKey"`
	Name     string    `gorm:"column:name"`
	FullName string    `gorm:"column:full_name"`
	Contact  uint64    `gorm:"column:contact"`
	Email    string    `gorm:"column:email"`
	Address  string    `gorm:"column:address"`
	Gender   string    `gorm:"column:Gender"`
	Password string    `gorm:"column:Password"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}
