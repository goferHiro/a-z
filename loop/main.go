package main

import "fmt"

func main(){
  days := []string {"M","T","W"}
  for i:=0;i<len(days);i++ {
    fmt.Println(days[i])
  }
  for y:= range days {
    fmt.Println(days[y])
  }
  for index,val := range days {
    fmt.Printf("%v %v\n",index,val)
  }

   //While loop
  i:=0
  for i<len(days){
    fmt.Println(i)
    i++
  }
    
}