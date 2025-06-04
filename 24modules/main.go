package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Understanding mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/contact",serveContact).Methods("GET")
	fmt.Println("http://localhost:1234")
	fmt.Println("http://localhost:1234/contact")

	log.Fatal(http.ListenAndServe(":1234", r))
}

func greeter() {
	fmt.Println("hey there mod users")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from vaibhav kothari on the web </h1>"))
}

func serveContact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Contact Page of the website </h1>"))
	w.Write([]byte("<h2>Hello world Page of the website </h2>"))
}
