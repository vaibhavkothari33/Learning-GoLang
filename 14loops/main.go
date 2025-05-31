package main
import "fmt"
func main() {
	fmt.Println("loops in go lang")

	days:= [] string{"Sunday","Tuesday","Wednesday","Friday","Saturday"}

	fmt.Println(days)

	// for i:=0; i<len(days);i++{
	// 	fmt.Println(days[i],i)
	// }

	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	// for i,day:= range days{
	// 	fmt.Printf("index is %v and value is %v\n",i,day)
	// }

	value:=1

	for value<10{
		if value == 5{
			break
		}
		fmt.Println("Value is ",value)
		value++
	}
}