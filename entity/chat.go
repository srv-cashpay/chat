package entity

import (
	"time"

	"gorm.io/gorm"
)

type ChatWidget struct {
	ID           string         `gorm:"primary_key,omitempty" json:"id"`
	FullName     string         `gorm:"full_name" json:"full_name"`
	Email        string         `gorm:"email;uniqueIndex" json:"email"`
	Whatsapp     string         `gorm:"whatsapp" json:"whatsapp"`
	BusinessName string         `gorm:"business_name" json:"business_name"`
	Description  string         `gorm:"description" json:"description"`
	CreatedBy    string         `gorm:"created_by" json:"created_by"`
	UpdatedBy    string         `gorm:"updated_by" json:"updated_by"`
	DeletedBy    string         `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
