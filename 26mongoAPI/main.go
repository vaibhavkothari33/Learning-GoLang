package main

import (
	"fmt"
	"net/http"

	"github.com/vaibhavkothari33/mongoapi/router"
)

func main() {
	fmt.Println("Mongo db api")
	r:= router.Router()
	fmt.Println("Server is getting started ...")
	http.ListenAndServe(":4000",r)
	fmt.Println("Listening at port number 4000 ..." )
}

// mongodb+srv://vaibhavkothari50:iZqTxT76IgSuoh9k@gobackend.jbas9pp.mongodb.net/
