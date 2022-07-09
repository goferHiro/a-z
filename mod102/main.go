package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
	"time"
  "log"
)

type Course struct {
	Id     string  `json:"courseId"`
	Name   string  `json:"courseName"`
	Author *Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
}

var courses []Course //fake db

func (c *Course) IsEmpty() bool {
	return c.Id == "" && c.Name == ""
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("<h1>All courses</h1>"))
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) //grab <id >
	for _, c := range courses {
		if c.Id == params["id"] {
			json.NewEncoder(w).Encode(c)
			return
		}
		json.NewEncoder(w).Encode("No courses")
		return
	}

}
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Data??")
		return
	}
	var sampleCourse Course

	json.NewDecoder(r.Body).Decode(&sampleCourse)
	if sampleCourse.IsEmpty() {
		json.NewEncoder(w).Encode("Data??")
		return
	}

	rand.Seed(time.Now().UnixNano())
	sampleCourse.Id = strconv.Itoa(rand.Int())
	courses = append(courses, sampleCourse)
	json.NewEncoder(w).Encode("Success")
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, c := range courses {
		if c.Id == params["id"] {
			courses = append(courses[:i], courses[i+1:]...) //delete course
			var sampleCourse Course
			_ = json.NewDecoder(r.Body).Decode(&sampleCourse)
			sampleCourse.Id = params["id"]
			courses = append(courses, sampleCourse)
			json.NewEncoder(w).Encode(sampleCourse)
		}
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DElete")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, c := range courses {
		if c.Id == params["id"] {
			courses = append(courses[:i], courses[i+1:]...) //delete course
			json.NewEncoder(w).Encode("Success")
			break
		}
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hey</h1>"))
}

func main() {
	fmt.Println("API")

	r := mux.NewRouter()

	//seedind

	courses = append(courses, Course{Id: "2", Name: "React", Author: &Author{Name: "Hiro"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getCourse).Methods("GET")

	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/courses", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))
}
