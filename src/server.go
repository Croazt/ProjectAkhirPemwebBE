package main

import (
	"ProjectAkhirPemweb/backend/src/routes"
	"ProjectAkhirPemweb/backend/src/utils"
	"ProjectAkhirPemweb/backend/src/utils/dbConfig"
	"log"
	"net/http"
	"os"
	"time"
)

type inisiationVar struct {
	l *log.Logger
}

func main() {
	dbConfig.ConnectDatabase()
	init := &inisiationVar{log.New(os.Stdout, "project-pemweb", log.LstdFlags)}
	serveMux := routes.Router(init.l)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	utils.GraceFullShutdown(init.l, s)
}
