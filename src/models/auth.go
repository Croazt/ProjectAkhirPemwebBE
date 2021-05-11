package models

import (
	"encoding/json"
	"io"

	"github.com/Croazt/ProjectAkhirPemwebBE/src/utils/dbConfig/dbModel"

	"github.com/go-playground/validator"
)

type UserLogin dbModel.UserLogin
type token struct {
	success string
	data    interface{}
}

func (ul *UserLogin) Validate() error {
	validate := validator.New()
	// validate.RegisterValidation("MEMEK", validateMemek)
	return validate.Struct(ul)
}

type UserLogins []*UserLogin

func (ul *UserLogins) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ul)
}

func (ul *UserLogin) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ul)
}

func TokenToJSON(w io.Writer, data interface{}) error {
	e := json.NewEncoder(w)
	return e.Encode(data)
}

func TokenFromJSON(r io.Reader, data interface{}) error {
	e := json.NewDecoder(r)
	return e.Decode(data)
}
