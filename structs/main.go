package main

import "fmt"


type User struct {
  Name string
  Age int
}

func main(){
  me := User{}
  me = User {"Hiro",22}
  fmt.Printf("Details %+v are ", me)
  fmt.Println(me.Age)
}