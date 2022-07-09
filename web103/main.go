package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const myurl = "https://dashboard-employee.herokuapp.com/employee"
const url = "https://dashboard-employee.herokuapp.com/employee"

func try(err error) {
	if err != nil {
		panic(err)
	}
}

func Post(url string) {
	reqBody := strings.NewReader(`
    {
    "name":"Hi",
    "designation":"Goofer"
    }
`)
  res, err:= http.Post(url,"application/json",reqBody)
  try(err)
  defer res.Body.Close()
  cnt, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(cnt))
}

func PostForm(url string){ //raw form data
  data := url.Values{}
  data.Add("name","niro")
  data.Add("designation","backend")
  _,_:= http.PostForm(url,data)
}  

func Get(url string) {
	res, err := http.Get(url)
	try(err)
	defer res.Body.Close()
	fmt.Println("Status Code ", res.StatusCode)

	cnt, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(cnt))

	//Alternate approach cnt=>string
	var resString strings.Builder
	byteCount, _ := resString.Write(cnt)
	fmt.Println(byteCount, resString.String())
}

func main() {
	//Get(myurl)
  //Post(url)
}
