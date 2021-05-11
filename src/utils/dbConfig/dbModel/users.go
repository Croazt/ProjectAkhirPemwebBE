package dbModel

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey, _"`
	Name      string    `json:"name" validate:"required" gorm:"<-"`
	BirthDay  uint8     `json:"birth_day" validate:"required" gorm:"<-:create"`
	Sex       bool      `json:"sex" validate:"required" gorm:"<-:create"`
	Phone     string    `json:"phone" validate:"required" gorm:"<-"`
	Email     string    `json:"email" validate:"required,email" gorm:"<-"`
	Password  string    `json:"password" `
	CreatedOn time.Time `json:"_" gorm:"<-:create"`
	UpdatedOn time.Time `json:"_" gorm:"<-:update"`
	DeletedOn time.Time `json:"_" gorm:"index"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
