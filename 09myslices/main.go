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

	// how to remove the value from slices

	var courses = []string{"Reactjs","JavaScript","Swift","Python","Ruby"}
	fmt.Println("The courses are" ,courses)
	var index int =2
	courses1 := append(courses[:index],courses[index+1:]...)
	courses = append(courses[:2],courses[3:]...)  // [Reactjs JavaScript Python Ruby]
	// courses = append(courses[:3])
	fmt.Println(courses)
	fmt.Println(courses1)
}