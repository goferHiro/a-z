package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
  "log"
)

func main() {
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	//http.Handle("/",serveHome)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Griit")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hey</h1>"))

}
