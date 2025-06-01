package main

import (
	"fmt"
	"io"
	"net/http"
	// "string"
)

const url = "https://vaibhavkothari.me"

func main() {
	fmt.Println("Web request")
	fmt.Println("the url is: ",url)

	responce , err:=http.Get(url)

	if err !=nil{
		panic(err)
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("responce ",responce)
	fmt.Println("")
	fmt.Println("")
	fmt.Printf("responce is of type %T ",responce)

	defer responce.Body.Close() // always close the connection

	body,err:=io.ReadAll(responce.Body)
	if err !=nil{
		panic(err)
	}
	fmt.Println("body is ",string(body))
}