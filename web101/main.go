package main

//handling web request

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func try(err error) {
	if err != nil {
		panic(err)
	}
}

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Rerquest")
	res, err := http.Get(url)

	fmt.Printf("Response %v %T", res, res)
	defer res.Body.Close()
	databytes, err := ioutil.ReadAll(res.Body)
	try(err)
	content := string(databytes)
	fmt.Println(content)
}
