package main

import  (
  "fmt"
  "strconv"
  "strings"
)

func main(){
  ageInString:="21 "
  age, err:= strconv.ParseFloat(strings.TrimSpace(ageInString),64)  
  if err!=nil{
    panic(err)
  }
  fmt.Printf("age is of %T",age)
}