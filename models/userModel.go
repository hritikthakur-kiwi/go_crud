package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	FullName string `gorm:"column:full_name"`
	Contact  uint64 `gorm:"column:contact"`
	Email    string `gorm:"column:email"`
	Address  string `gorm:"column:address"`
	Gender   string `gorm:"column:Gender"`
	Password string `gorm:"column:Password"`
	
}
