package routes

import (
	"ProjectAkhirPemweb/backend/src/handlers"
	"ProjectAkhirPemweb/backend/src/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func userRouter(serveMux *mux.Router, l *log.Logger, handler http.Handler) {

	userHandler := handlers.NewUser(l)
	authHandler := handlers.NewAuth(l)
	userCreateRouter := serveMux.PathPrefix("/users").Subrouter()
	userCreateRouter.HandleFunc("/", userHandler.AddUser).Methods("POST")
	userCreateRouter.Handle("/{ID:[0-9]+}", middleware.MiddlewareUserValidation(http.HandlerFunc(userHandler.UpdateUser))).Methods("PUT")

	userRouter := serveMux.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", userHandler.GetUser).Methods("GET")
	userRouter.HandleFunc("/{ID:[0-9]+}", userHandler.GetOneUser).Methods("GET")
	userRouter.HandleFunc("/{ID:[0-9]+}", userHandler.DeleteUser).Methods("DELETE")
	userRouter.Handle("/register", middleware.MiddlewareUserValidation(http.HandlerFunc(authHandler.Register))).Methods("POST")
	userRouter.Handle("/login", middleware.MiddlewareLoginValidation(http.HandlerFunc(authHandler.Login))).Methods("POST")
}
