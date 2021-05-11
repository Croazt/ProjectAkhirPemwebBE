package utils

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func GraceFullShutdown(l *log.Logger, s *http.Server) {
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	l.Println("Received terminate, gracefull shutdown ", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
