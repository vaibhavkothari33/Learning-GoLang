package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("welcome to web server")
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000"

	responce,err:= http.Get(myurl)

	if err !=nil{
		panic(err)
	}

	
	defer responce.Body.Close()

	fmt.Println("Status Code :",responce.StatusCode)
	fmt.Println("Status :",responce.Status)
	fmt.Println("Content length :",responce.ContentLength)

	content,_:=io.ReadAll(responce.Body)
	fmt.Println(string(content))
	fmt.Println(content)
}