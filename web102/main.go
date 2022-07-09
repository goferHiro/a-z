package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payment&paymentid=1221gh"

func main() {
	fmt.Println("web 102 Handling URL")
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
  fmt.Println(result.Scheme)
  fmt.Println(result.Host)
  fmt.Println(result.Hostname())
  fmt.Println(result.Port())
  fmt.Println(result.RawQuery)

  qparams := result.Query()
  fmt.Printf("%T",qparams) //map

  for param,values:=range qparams{
    fmt.Println(param,values)
  }

  //Url ->string
  partsOfUrl := &url.URL{ //& imp
    Scheme:"https",
    Host:"lco.dev",
    Path:"/tutc",
    RawPath:"user=hiro", //qparms
  }
  customUrl := partsOfUrl.String()
  fmt.Println(customUrl )
}
