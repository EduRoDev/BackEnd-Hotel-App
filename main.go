package main

import (
	"log"
	"net/http"
	"os"

	"github.com/DxKaizer/Hotel_Backend/pkg/controllers"
	"github.com/DxKaizer/Hotel_Backend/pkg/model"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "Hotel-api ", log.LstdFlags)

	hl := controllers.NewLogger(l)
	mx := mux.NewRouter()
	model.Migraciones()

	mx.HandleFunc("/Get", hl.GetHotel).Methods("GET")
	mx.HandleFunc("/post", hl.PostHotel).Methods("POST")

	l.Println("Starting server on port 8080")
	l.Fatal(http.ListenAndServe(":8080", mx))

}
