package models

import "time"

type User struct {
	ID        string    `gorm:"primary_key" form:"id" json:"id"`
	FullName  string    `form:"full_name" json:"full_name" binding:"required"`
	Email     string    `form:"email" json:"email" binding:"required"`
	Password  string    `form:"password" json:"password" binding:"required"`
	Phone     string    `form:"phone" json:"phone" binding:"required"`
	Document  string    `form:"document" json:"document" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" form:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" form:"update_at" json:"update_at"`
}
