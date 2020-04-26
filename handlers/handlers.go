package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/solrac97gr/twitter-fake/middlewares"
	"github.com/solrac97gr/twitter-fake/routers"
)

/*Handlers : set the port,cors,handlers and then serve the api*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.CheckDB(routers.Register).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Println("Server run in http://localhost:" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
