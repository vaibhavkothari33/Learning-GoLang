package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json")

	myCourses := []course{
		{"Reactjs", 299, "Vaibhavkothari.me", "abc123", []string{"web-dev", "javascript"}},
		{"JavaScript", 199, "Vaibhavkothari.me", "dsfbrth5y5", []string{"app-dev", "typescript"}},
		{"NextJs", 399, "Vaibhavkothari.me", "sdfdsf3r3", nil},
	}

	// package this data as json data

	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
	DecodeJson()

}
func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		  "coursename": "Reactjs",
                "Price": 299,
                "website": "Vaibhavkothari.me",
                "tags": [
                        "web-dev",
                        "javascript"
                ]
	}`)

	var mycourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &mycourse)
		fmt.Printf("%#v\n",mycourse)
	}else{
		fmt.Println("Json is not valid")
	}

	// some cases where we need to add the data in key value pair

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb,&myOnlineData)
	fmt.Printf("%#v\n",myOnlineData)

	for k,v := range myOnlineData{
		fmt.Printf("Key is %v and value is %v and type is: %T \n",k,v,v)
	}
}


// Welcome to json
// [
//         {
//                 "coursename": "Reactjs",
//                 "Price": 299,
//                 "website": "Vaibhavkothari.me",
//                 "tags": [
//                         "web-dev",
//                         "javascript"
//                 ]
//         },
//         {
//                 "coursename": "JavaScript",
//                 "Price": 199,
//                 "website": "Vaibhavkothari.me",
//                 "tags": [
//                         "app-dev",
//                         "typescript"
//                 ]
//         },
//         {
//                 "coursename": "NextJs",
//                 "Price": 399,
//                 "website": "Vaibhavkothari.me"
//         }
// ]
