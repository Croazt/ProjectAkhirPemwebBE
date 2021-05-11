package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID int
	Sport  string
	Date   time.Time
	Title  string
	Time   string
	Place  string
}
