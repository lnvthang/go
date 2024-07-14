package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int       `gorm:"column:id;primaryKey"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Gmail     string    `gorm:"column:gmail"`
	Fullname  string    `gorm:"column:fullname"`
	Dob       time.Time `gorm:"column:dob"`
	Status    bool      `gorm:"column:status"`
	IsDeleted int       `gorm:"column:is_deleted"`
	CreatedBy int       `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedBy int       `gorm:"column:updated_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedBy int       `gorm:"column:deleted_by"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
