package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	handlerscustomers "github.com/DxKaizer/Hotel_Backend/pkg/controllers/HandlersCustomers"
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
	cl := handlerscustomers.NewLogger(l)

	//rutas para los clientes
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

	//rutas para el usuario del hotel
	mx.HandleFunc("/Customer/get", cl.GetCustomers).Methods("GET")
	mx.HandleFunc("/Customer/get/{id:[0-9]+}", cl.GetCustomersID).Methods("GET")
	mx.HandleFunc("/Customer/post", cl.PostCustomers).Methods("POST")
	//mx.HandleFunc("/Customer/put/{id:[0-9]+}", cl.PutCustomers).Methods("PUT")
	//mx.HandleFunc("/Customer/delete/{id:[0-9]+}", cl.DeleteCustomers).Methods("DELETE")

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mx,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Println("Starting server on port :8080")
		err := srv.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	l.Println("Received terminated, graceful shutdown", sig)
	tc, err := context.WithTimeout(context.Background(), 30*time.Second)

	if err != nil {
		l.Println(err)
	}
	srv.Shutdown(tc)

}
