package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"math/rand"  // Add this import
	"github.com/gorilla/mux"
)

// model  for course - file

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}


type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}


// face db

var courses []Course
// middleware , helper - file

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}


func main(){
	fmt.Println("API  - GO")
	r:= mux.NewRouter()


	// data seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJs", CoursePrice: 299, Author: &Author{Fullname: "Hitesh C", Website: "lfd.co"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "ReactJs", CoursePrice: 299, Author: &Author{Fullname: "Hitesh C", Website: "lfd.co"}})



	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))


}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API using Go language</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}


func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)

	// loop through courses, find matching id and return the reponse

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course 
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}


	// generate unique id, string
	// append  course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	return
}

// func updateOneCourse (w http.ResponseWriter, r *http.Request){
// 	fmt.Println("create one course")
// 	w.Header().Set("Content-Type", "application/json")

// 	// first - grap id from req

// 	params := mux.Vars(r)

// 	// loop, once we hit the id and then remove, add  with my ID 

// 	for index, course := range courses {
// 		if course.CourseId == params["id"] {
// 			courses = append(courses[:index], courses[index+1:]...)
// 			var course Course
// 			_ = json.NewDecoder(r.Body).Decode(&course)
// 			course.CourseId = params["id"]
// 			courses = append(courses, course)
// 			json.NewEncoder(w).Encode(course)
// 			return
// 			// TODO : send a repsone when id is not found
// 		} else {
// 			json.NewEncoder(w).Encode("Sorry your ID is not found try again")
// 		}
// 	}


// }

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from req
	params := mux.Vars(r)

	// loop to find the course with the specified id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// remove the old course with the specified id
			courses = append(courses[:index], courses[index+1:]...)

			// decode the request body to get the new course data
			var newCourse Course
			err := json.NewDecoder(r.Body).Decode(&newCourse)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// set the id of the new course
			newCourse.CourseId = params["id"]

			// append the new course to the slice
			courses = append(courses, newCourse)

			// encode and send the new course as the response
			json.NewEncoder(w).Encode(newCourse)
			return
		}
	}

	// if the loop completes without finding the course, send an error response
	json.NewEncoder(w).Encode("Sorry, the ID is not found. Please try again.")
}


func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id remove, (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

}
