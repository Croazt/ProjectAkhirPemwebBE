package models

import (
	"encoding/json"
	"io"
	"log"

	"github.com/Croazt/ProjectAkhirPemwebBE/src/utils/dbConfig"
	"github.com/Croazt/ProjectAkhirPemwebBE/src/utils/dbConfig/dbModel"

	"github.com/go-playground/validator"
)

type User dbModel.User

func (u *User) Validate() error {
	validate := validator.New()
	// validate.RegisterValidation("MEMEK", validateMemek)
	return validate.Struct(u)
}

func AddUser(u *User) bool {
	if err := dbConfig.DB.Create(u).Error; err == nil {
		return false
	} else {
		return true
	}
}

func UpdateUser(u *User, id uint) bool {
	user := GetOneUser(id)

	if user == nil {
		return false
	}

	u.ID = id
	if err := dbConfig.DB.Save(u); err == nil {
		return false
	}

	return true
}

func DeleteUser(id uint) bool {
	user := GetOneUser(id)

	if user == nil {
		return false
	}

	if err := dbConfig.DB.Delete(&User{}, id); err == nil {
		return false
	}

	return true
}

type Users []*User

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func GetUser() Users {
	var users []*User
	if err := dbConfig.DB.Find(&users).Error; err != nil {
		log.Println("error", err)
		return nil
	}

	return users

}

func GetUserByEmail(email string) *User {
	var users *User
	if err := dbConfig.DB.Find(&users, "email = ?", email).Error; err != nil {
		log.Println("error", err)
		return nil
	}

	return users

}

func GetOneUser(id uint) *Users {
	var users *Users
	if err := dbConfig.DB.First(&users, id).Error; err != nil {
		return nil
	}

	return users

}
