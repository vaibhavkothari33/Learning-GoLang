package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://bennett.edu:3000/learn?coursename=btech&paymentid=skdjgjdhfkjghkfjdhdjkfhskhfkjshuyityi"

// const myurl string = "https://bennett.edu.in"

func main() {
	fmt.Println("urls in golang")
	fmt.Println("the url is ", myurl)

	// parsing  the url
	result, err := url.Parse(myurl) // url ko parse karo

	// fmt.Println("protocal : ",result.Scheme)
	// fmt.Println("host id/number :",result.Host)
	// fmt.Println("Port number :",result.Port())
	// fmt.Println("Path :",result.Path)
	// fmt.Println("route :",result.RawQuery)
	// fmt.Println("params :",result.RawPath)

// 	protocal :  http
// 	host id/number : bennett.edu:3000
// 	Port number : 3000
// 	Path : /learn
// 	route : coursename=btech&paymentid=skdjgjdhfkjghkfjdhdjkfhskhfkjshuyityi
// 	params :

	qparams := result.Query()
	fmt.Printf("the type of query params are: %T",qparams)
	fmt.Println("")
	fmt.Println(qparams["coursename"])


	for _,val := range qparams{
		fmt.Println("Params is ",val)
	}
	if err != nil {
		panic(err)
	}

	fmt.Println("the responce is ", result)

	partsOfurl:= &url.URL{Scheme: "https",Host:"vaibhavkothari.me",Path:"/resume.pdf",RawPath: "user=vaibhav"}

	anotherUrl:= partsOfurl.String()
	fmt.Println(anotherUrl)
}
