package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Croazt/ProjectAkhirPemwebBE/src/models"
)

func MiddlewareUserValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := models.User{}
		err := user.FromJSON(r.Body)
		if err != nil {
			log.Println("[ERROR] dezerializing product", err)
			http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		err = user.Validate()
		if err != nil {
			log.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprint("Error vlidating product : %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, user)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}

func MiddlewareLoginValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := models.UserLogin{}
		err := user.FromJSON(r.Body)
		if err != nil {
			log.Println("[ERROR] dezerializing product", err)
			http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		err = user.Validate()
		if err != nil {
			log.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprint("Error vlidating product : %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, user)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
