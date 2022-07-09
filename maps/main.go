package main

import "fmt"

func main(){
  maps := make(map[string]string)
  maps["FR"] = "France"
  fmt.Println("maps are ", maps)
  fmt.Println("maps for FR ", maps["FR"])
  
  for key,val:=range maps{
    fmt.Printf("For key %v, value is %v",key,val)
  }
  
  delete(maps,"FR")
  
}