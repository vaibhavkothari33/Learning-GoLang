package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/vaibhavkothari33/mongoapi/router"
)

func main() {
	fmt.Println("Mongo db api")

	r := router.Router()

	// Set up CORS options
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(r)

	fmt.Println("Server is getting started ...")
	log.Fatal(http.ListenAndServe(":4000", handler))
}
