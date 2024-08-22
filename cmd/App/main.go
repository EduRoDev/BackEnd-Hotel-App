package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DxKaizer/Hotel_Backend/pkg/controllers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "Hotel-api ", log.LstdFlags)
	hl := controllers.NewLogger(l)
	mx := mux.NewRouter()
	mx.Handle("/", hl)
	s := &http.Server{
		Addr:         ":8080",
		Handler:      mx,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	func() {
		l.Println("Starting server on port 8080")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
}
