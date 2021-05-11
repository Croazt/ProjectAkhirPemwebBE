package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Croazt/ProjectAkhirPemwebBE/src/middleware"
	"github.com/Croazt/ProjectAkhirPemwebBE/src/models"

	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	l *log.Logger
}

func NewAuth(l *log.Logger) *Auth {
	return &Auth{l}
}

type MyClaims struct {
	jwt.StandardClaims
	Nama  string `json:"Nama"`
	Email string `json:"Email"`
}

func (a *Auth) Login(rw http.ResponseWriter, r *http.Request) {
	data := r.Context().Value(middleware.KeyProduct{}).(models.UserLogin)

	user := models.GetUserByEmail(data.Email)
	a.l.Println(user)
	if user == nil {
		http.Error(rw, "Cannot find user with email given", http.StatusNotFound)
		return
	}

	if user.Password != data.Password {
		http.Error(rw, "Password not matched", http.StatusNotAcceptable)
		return
	}

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Battlesport",
			ExpiresAt: time.Now().Add(9000 * time.Hour).Unix(),
		},
		Nama:  user.Name,
		Email: user.Email,
	}

	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	).SignedString([]byte("128340asdnkadasd"))

	if err != nil {
		http.Error(rw, "Cannot sign token", http.StatusNotFound)
		return
	}
	models.TokenToJSON(rw, token)

}

func (a *Auth) Register(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.KeyProduct{}).(models.User)

	if err := models.AddUser(&user); err {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Battlesport",
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
		},
		Nama:  user.Name,
		Email: user.Email,
	}

	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	).SignedString([]byte("128340asdnkadasd"))

	if err != nil {
		http.Error(rw, "Cannot sign token", http.StatusNotFound)
		return
	}
	models.TokenToJSON(rw, token)

}
