package model

import (
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid-ossp.gen_random_uuid();primaryKey" json:"id"`
	Name     string    `gorm:"column:name"`
	FullName string    `gorm:"column:full_name"`
	Contact  uint64    `gorm:"column:contact"`
	Email    string    `gorm:"column:email"`
	Address  string    `gorm:"column:address"`
	Gender   string    `gorm:"column:Gender"`
	Password string    `gorm:"column:Password"`
	Posts    []Post    `gorm:"foreignKey:UserId"`
}
