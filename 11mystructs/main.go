package main

import "fmt"

func main() {
	fmt.Println("This is structs in go")

	vaibhav:=User{"Vaibhav","Vaibhavkothari@go.dev",true,20}
	fmt.Println("detail",vaibhav)
	fmt.Printf("detail %T",vaibhav)

}

type User struct {
	Name  string
	Email string
	Status bool
	Age int
}
