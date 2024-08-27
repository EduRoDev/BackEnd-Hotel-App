package main

import (
	"log"
	"net/http"
	"os"

	controllers "github.com/DxKaizer/Hotel_Backend/pkg/controllers/HandlersHotel"
	handlersroom "github.com/DxKaizer/Hotel_Backend/pkg/controllers/HandlersRoom"
	"github.com/DxKaizer/Hotel_Backend/pkg/database"

	"github.com/gorilla/mux"
)

func main() {
	database.Migraciones()
	l := log.New(os.Stdout, "Hotel-api ", log.LstdFlags)

	hl := controllers.NewLogger(l)
	rl := handlersroom.NewLogger(l)
	mx := mux.NewRouter()

	//rutas para el Hotel
	mx.HandleFunc("/get", hl.GetHotel).Methods("GET")
	mx.HandleFunc("/get/{id:[0-9]+}", hl.GetHotelID).Methods("GET")
	mx.HandleFunc("/post", hl.PostHotel).Methods("POST")
	mx.HandleFunc("/put/{id:[0-9]+}", hl.PutHotels).Methods("PUT")
	mx.HandleFunc("/delete/{id:[0-9]+}", hl.DeleteHotel).Methods("DELETE")

	//rutas para el cuarto de los hoteles
	mx.HandleFunc("/Room/get", rl.GetRoom).Methods("GET")
	mx.HandleFunc("/Room/get/{id:[0-9]+}", rl.GetRoomID).Methods("GET")
	mx.HandleFunc("/Room/post", rl.PostRoom).Methods("POST")
	mx.HandleFunc("/Room/put/{id:[0-9]+}", rl.PutRoom).Methods("PUT")
	mx.HandleFunc("/Room/delete/{id:[0-9]+}", rl.DeleteRoom).Methods("DELETE")

	l.Println("Starting server on port :8080")
	l.Fatal(http.ListenAndServe(":8080", mx))

}
