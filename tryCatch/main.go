package main

import (
  "bufio"
  "os"
  "fmt"
)

func main(){
  reader:= bufio.NewReader(os.Stdin)
  fmt.Println("Enter something")
  input,err := reader.ReadString('\n')
  if err!=nil{fmt.Println("There is an err ",err)
  }else{
   fmt.Println("Thanks for the i/p ",input)
  }
}