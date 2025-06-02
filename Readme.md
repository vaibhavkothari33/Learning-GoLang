# Boiler plate for go is pkmg

## how to take the input

```go
package main

import "fmt"

func main(){
    fmt.Println("Enter your name")
    reader:=bufio.NewReader(os.Stdin)

    input,_:=reader.ReadString('\n')
    fmt.Println("your name is ",input)

}
```

## how to do the conversion

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main(){
	fmt.Println("enter the number between 1 and 5")

	reader:=bufio.NewReader(os.Stdin)

	input,_:=reader.ReadString('\n')
	fmt.Println("the number is ",input)

	// addition,err:=strconv.ParseFloat(strings.TrimSpace(input),64)

	addition,err:=strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println("the addition is ",addition +5)
	}

}
```

## time control

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println("Formated Time")
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate:=time.Date(2020,time.January,19,23,23,0,0,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	// date time week
}
```

# Memory allocation and deallocation happes automatically in go lang

```go
new ()

allocate memory but no init
you will get a memory address

zeroes storage
```

```go
make()
Allocate memory and init
you will get a memory address
non zeroes storage
```

garbage collection happed automatically
out of scope or nil

## Pointers in Go Lang

passing the memory address not the variable
actual value reference
& ===> reference

```go
package main

import "fmt"

func main() {
	fmt.Println("Pointers in go lang")

	// var ptr* int
	// fmt.Println("The value of pointer is",ptr)

	myNumber:= 23

	var ptr = &myNumber
	fmt.Println("the value of ptr is ",ptr) // )xc0000008798
	fmt.Println("the actual value of ptr is ",*ptr) // 23
	*ptr = *ptr *2
	fmt.Println("the new value is",*ptr) // 46
}
```

## arrays in Golang no need to loop just print for direct array

```go

package main

import "fmt"

func main() {
	fmt.Println("The arrays in go lang")

	var fruitlist [4] string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Mango"
	// fruitlist[2] = ""
	fruitlist[3] = "Banana"

	fmt.Println("Fruit list is ", fruitlist) // Fruit list is  [Apple Mango (space) Banana]

	fmt.Println("The length of the fruitlist is ",len(fruitlist)) // 4 but i only have 3 items

	var vegList = [3] string{"Potato","Beans","Onion"}

	fmt.Println("Veg List is ",len(vegList))
	fmt.Println("Veg List is ",vegList)
	// fmt.Println("Veg List is ",vegList[5]) // out of bound
	fmt.Println("Veg List is ",vegList[2])
}
```

## Slices in go lang (more powerfull then array)

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("The use of slices in go lang")

	var fruitlist = [] string{"Apple", "Banana","Grapes"}
	fmt.Printf("Type of fruitlist is %T\n",fruitlist)
	fmt.Println("The fruit list is ",fruitlist)

	fruitlist = append(fruitlist, "Mango","Chicku") //same python what to  add
	fmt.Println("the new fruitlist is ",fruitlist)

	// fruitlist = append(fruitlist[1:])
	fruitlist =append(fruitlist[:])
	// fruitlist =append(fruitlist[:3])
	// fruitlist = append(fruitlist[1:3]) // print 1 and 2 nd value
	fmt.Println("the new list is",fruitlist)

	// new() and make()

	highScores:= make([]int,4)
	highScores[0] = 234
	highScores[1] = 976
	highScores[2] = 123
	highScores[3] = 567

	highScores  = append(highScores, 555,666,321) // [234 976 123 45 555 666 321]
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)

	fmt.Println("the sorted list is ",highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
}
```

## for each

```go
	for key, value:=range languages{
		fmt.Println(key,value)
	}
```

## maps in go lang

```go
package main

import "fmt"

func main() {
	fmt.Println("Maps in go lang")

	languages:=make(map[string]string)

	languages["JS"] = "Javascript"
	languages["GO"] = "GOLANG"
	languages["PY"] = "Python"

	fmt.Println("List of the all the languages",languages)
	fmt.Println("JS shorts for: ", languages["JS"])
	fmt.Println("PY shorts for: ", languages["PY"])

	delete(languages,"GO")
	fmt.Println("List of all languages: ",languages)
	}
```

## there is no interiance in go lang and classes no super parent

# functions in go lang

✅ Public (Exported):
Function name starts with an uppercase letter

Accessible from other packages

```go
// Public function
func PrintMessage() {
    fmt.Println("This is a public function")
}
```

❌ Private (Unexported):
Function name starts with a lowercase letter

Accessible only within the same package

```go
// Private function
func printMessage() {
    fmt.Println("This is a private function")
}

```

```go
package main
import "fmt"

func greater(){
	fmt.Println("Hello world")
	result:=adder(3,5)
	fmt.Println("the sum is",result)

	proResult:= proAdder(2,5,4,8,8)
	fmt.Println(proResult)
}

func adder(valOne int,valTwo int) int{
	return valOne + valTwo
}

func proAdder(values ... int)int{
	total := 0
	for _,value:= range values{
		total+=value
	}
	return total
}
func main() {
	fmt.Println("Welcome to entry point in go lang")
	greater()
}
```

## loops

```go
for i,day:= range days{
	fmt.Printf("index is %v and value is %v\n",i,day)
}
```

# parts of url

```go
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
```
