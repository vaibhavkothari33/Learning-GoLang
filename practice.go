package main

import (
	"fmt"
	"io"
	"net/http"
	// "strings"
)
const url = "https://bennett.edu.in"
func main() {
	fmt.Println("Url in golang")

	fmt.Println("the url is ",url)

	responce,err:= http.Get(url)

	if err!= nil{
		panic(err)
	}
	fmt.Println("the responsce is: ",responce)

	defer responce.Body.Close()

	body,_:=io.ReadAll(responce.Body)

	fmt.Println("the body is ",body)
	fmt.Println("the body is ",string(body))

}