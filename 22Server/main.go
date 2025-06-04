package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to web server")
	PerformGetRequest()
	PerformPostJsonrequest()
	PerformFormPostRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"

	responce,err:= http.Get(myurl)

	if err !=nil{
		panic(err)
	}

	
	defer responce.Body.Close()

	fmt.Println("Status Code :",responce.StatusCode)
	fmt.Println("Status :",responce.Status)
	fmt.Println("Content length :",responce.ContentLength)

	// content,_:=io.ReadAll(responce.Body)
	// fmt.Println(string(content))
	// fmt.Println(content)

	var responceString strings.Builder
	content ,err := io.ReadAll(responce.Body)

	if err !=nil{
		panic(err)
	}

	byteCount ,_:= responceString.Write(content)

	fmt.Println("")
	fmt.Println("")
	fmt.Println(byteCount)
	fmt.Println(responceString.String())

}


func PerformPostJsonrequest(){
	// send post request 
	const myurl = "http://localhost:8000/post"


	requestBody := strings.NewReader(`
	{
		"username" : "Vaibhav",
		"age" : 19,
		"address" : "Jaipur"
	}`)

	responce ,err :=http.Post(myurl, "application/json",requestBody)

	if err != nil{
		panic(err)
	}

	defer responce.Body.Close()

	content, _:= io.ReadAll(responce.Body)

	fmt.Println("")
	fmt.Println("")
	fmt.Println(string(content))
	// fmt.Println("responce",)

}

func PerformFormPostRequest(){

	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname","vaibhav")
	data.Add("lastname","kothari")
	data.Add("age","20")

	responce,err :=http.PostForm(myurl,data)

	if err!= nil {
		panic(err)
	}

	defer responce.Body.Close()

	content , _:= io.ReadAll(responce.Body)
	fmt.Println("")
	fmt.Println("")
	fmt.Println(string(content))

}