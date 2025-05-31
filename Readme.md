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

## there is no interiance in go lang and classes  no super parent