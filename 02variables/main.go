package main

import "fmt"

const LoginToken string = "qwertyuio" // first aph need to be capital to  -> make it public 

func main(){
	fmt.Println("--------------------------")
	fmt.Println("hello world")
	var username string = "vaibhav";
	var isloggedin bool = true
	var smallVal uint8 = 255
	var bigVal uint8 = 255
	var anotherVariable int;
	var emptyString string;
	fmt.Println(username)
	fmt.Printf("variable is of type %T \n",username)
	fmt.Printf("variable is of type %T \n",isloggedin)
	fmt.Printf("variable is of type %T \n",smallVal)
	fmt.Printf("variable is of type %T \n",bigVal)
	fmt.Printf("variable is of type %T \n",anotherVariable)
	fmt.Println("variable is of type",anotherVariable)
	fmt.Println("variable is of type",emptyString)

	var website = "hello world"
	fmt.Println(website)
	website = "hello"
	fmt.Println(website)
	name:="vaibhav kothari"
	fmt.Println(name,username)
	age:=20

	fmt.Println(age)
	 
	fmt.Println("-------------------------------------------")

}
