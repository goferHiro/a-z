package main

import "fmt"
func main(){
  var fruitList[4] string

  fruitList[0] = "Apple"
  fruitList[1] = "Peach"
  fruitList[2] = "Mango"

  fmt.Println("Array ", fruitList)
  fmt.Println("Length ", len(fruitList)) //4

  var vegList = [3]string {"potato"}
  fmt.Println("VegList ", vegList)
  
}