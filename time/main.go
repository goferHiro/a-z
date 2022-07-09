package main

import(
  "time"
  "fmt"
)
func main(){
  now := time.Now()
  fmt.Println("Time is ",now)
  formattedNow := now.Format("01-02-2006 15:04:05 Monday")//std time Jan
  fmt.Println("Formated time",formattedNow)
  newDate := time.Date(2022,time.April,10,23,23,04,0,time.UTC)
  fmt.Println(newDate)
}