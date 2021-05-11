package models

import (
	"gorm.io/gorm"
)

type SportPlace struct {
	gorm.Model
	Name            string
	Rating          uint
	Address         string
	OperationalHour string
	Description     string
}
