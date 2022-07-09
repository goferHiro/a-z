package main

import ("fmt"
        "sort")

func main(){
  fruits := [] string{"Apple"}
  fruits = append(fruits,"Mango")
  fruits = append(fruits[1:],"Apricot") //starts at 1 so no apple
  
  fmt.Println(fruits) //apple deleted
  
  scores :=  make([] int ,1)

  scores[0] = 1
  //scores[1] =2 //index out of range
  scores = append(scores,2)
  sort.Ints(scores)
}