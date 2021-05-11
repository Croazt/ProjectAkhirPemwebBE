package models

import (
	"time"

	"gorm.io/gorm"
)

type Sport struct {
	gorm.Model
	UserID      int
	Sport       string
	Date        time.Time
	Title       string
	Description string
}
