package controllers

import (
	"log"
	"net/http"
)

type Hotel struct {
	l *log.Logger
}

// ServeHTTP implements http.Handler.
func (h *Hotel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Get(w, r)
}

func NewLogger(l *log.Logger) *Hotel {
	return &Hotel{l}
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	rp := map[string]string{
		"message": "Hola mundo, como estas ?",
	}
	w.Write([]byte(rp["message"]))
}
