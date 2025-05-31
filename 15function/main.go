package main
import "fmt"

func greater(){
	fmt.Println("Hello world")
	result:=adder(3,5)
	fmt.Println("the sum is",result)

	proResult,mymessage:= proAdder(2,5,4,8,8)
	fmt.Println(proResult)
	fmt.Println(mymessage)
}

func adder(valOne int,valTwo int) int{
	return valOne + valTwo
}

func proAdder(values ... int)(int,string){
	total := 0
	for _,value:= range values{
		total+=value
	}
	return total,"hi"
}
func main() {
	fmt.Println("Welcome to entry point in go lang")
	greater()
}