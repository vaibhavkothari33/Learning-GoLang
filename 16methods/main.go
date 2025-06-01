package main
import "fmt"
type User struct{
	Name string
	Age int
}

func (u User) GetStatus(){
	fmt.Println("")
}

func (u User) NewMail(){
	fmt.Println("the new mail is vaibhav@gmail.com")
}
func main() {
	
}