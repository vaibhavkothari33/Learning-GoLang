package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for courses -file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` // type is a pointer
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake db

var courses []Course

// middlewares , helper = file

// func (c *Course) IsEmpty() bool {
// 	return c.CourseId == ""
// 	// return c.CourseId == "" && c.CourseName == ""
// }
func (c *Course) IsEmpty() bool {
	return c.CourseName == "" // or also check Author if needed
}


func main() {
	fmt.Println("Api Hello world")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "react js", CoursePrice: 299, Author: &Author{Fullname: "Vaibhav Kothari", Website: "vaibhavkothari.me"}})

	courses = append(courses, Course{CourseId: "3", CourseName: "Next js", CoursePrice: 199, Author: &Author{Fullname: "Nalin Thakur", Website: "vaibhavkothari.me"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "Tailwind css", CoursePrice: 499, Author: &Author{Fullname: "Gourav Kumar", Website: "go.dev"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")
	r.HandleFunc("/courses", deleteAllCourse).Methods("DELETE")

	// listen
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

//serveHome route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to api by go lang server </h1>"))
}

// seeding
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all the courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)

	fmt.Printf("%T", params)
	// loop through courses find matching id and return the responce

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		
	}

	// what about {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the json")
		return
	}

	//checking if the name exist already

	for _, checkname := range courses {
		if checkname.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course with same name already exist")
			return
		}
	}

	for _, checkId := range courses {
		if checkId.CourseId == course.CourseId {
			json.NewEncoder(w).Encode("Course with same id already exist")
			return
		}
	}
	//generate the unique id,string

	//append new courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updating the course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop , id ,remove ,add with my id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var updatedCourse Course // error   <-------------------->
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
			updatedCourse.CourseId = params["id"]
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return

		}
	}
	// send a responce when id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Delete course successfully")
			break
		}
	}

}
func deleteAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting all the course")
	w.Header().Set("Content-Type", "application/json")

	courses = []Course{} // reset the slice to an empty slice

	json.NewEncoder(w).Encode("All courses deleted successfully")

}
