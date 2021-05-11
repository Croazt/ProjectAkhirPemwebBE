package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router(l *log.Logger) http.Handler {

	serveMux := mux.NewRouter()

	userRouter(serveMux, l, serveMux)

	return serveMux
}
