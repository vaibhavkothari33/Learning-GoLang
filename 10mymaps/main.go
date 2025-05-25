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


	// for key,value:=range languages{
	// 	fmt.Printf("for key %v, value is %v\n", key,value)
	// }

	for key, value:=range languages{
		fmt.Println(key,value)
	}
}

