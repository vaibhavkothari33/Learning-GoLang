package main
import "fmt"
func main() {
	fmt.Println("if else in go lang")

	loginCount:= 23
	var result string

	if loginCount<2 {
		result = "register User"
	} else if loginCount>10{
		result = "Watch out"
	} else{
		result = "Somthig else"
	}

	fmt.Println(result)

	if num:= 3; num<10{
	fmt.Println("num is less then 10")
	} else{
		fmt.Println("num is not less than 10")
	}
}