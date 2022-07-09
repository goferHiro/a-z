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
  me.SetAge(21)
} 

func (u User) GetAge (){
  fmt.Println(u.Age)
}


func (u User) SetAge (age int){
  u.Age = age //Doesnot change original pass pointerinstead or return u.
  
}