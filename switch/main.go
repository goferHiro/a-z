package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main(){
  
  rand.Seed(time.Now().UnixNano())
  diceNumber := rand.Intn(6)+1  //DIce number random 
   fmt.Println("Dice no ", diceNumber)

  switch diceNumber {
    case 1: 
      fmt.Println("u r 1")    
    case 2: 
      fmt.Println("u r 2")    
    case 3: 
      fmt.Println("u r 3")    
    case 4: 
      fmt.Println("u r 4")   
      fallthrough //runs next case too
    case 5: 
      fmt.Println("u r 5")
    case 6: 
      fmt.Println("u r 6")
     default:
      fmt.Println("Default")
  }
  //Try in goplayground u get same value over
}