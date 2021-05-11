package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Croazt/ProjectAkhirPemwebBE/src/middleware"
	"github.com/Croazt/ProjectAkhirPemwebBE/src/models"

	"github.com/gorilla/mux"
)

type Users struct {
	l *log.Logger
}

func NewUser(l *log.Logger) *Users {
	return &Users{l}
}

func (u *Users) GetUser(rw http.ResponseWriter, r *http.Request) {
	listUser := models.GetUser()
	err := listUser.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (u *Users) GetOneUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["ID"])

	if err != nil {
		http.Error(rw, "Cannot convert id to num", http.StatusBadRequest)
		return
	}

	user := models.GetOneUser(uint(id))

	if user == nil {
		http.Error(rw, "User Not Found", http.StatusNotFound)
	}

	if err := user.ToJSON(rw); err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}

}

func (u *Users) AddUser(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.KeyProduct{}).(models.User)

	err := models.AddUser(&user)

	if err {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (u *Users) UpdateUser(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["ID"])

	if err != nil {
		http.Error(rw, "Cannot convert id to num", http.StatusBadRequest)
		return
	}

	user := r.Context().Value(middleware.KeyProduct{}).(models.User)

	if err := models.UpdateUser(&user, uint(id)); !err {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (u *Users) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["ID"])

	if err != nil {
		http.Error(rw, "Cannot convert id to num", http.StatusBadRequest)
		return
	}

	if err := models.DeleteUser(uint(id)); !err {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}
