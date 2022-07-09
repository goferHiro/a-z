package main

import "fmt"

func main(){

 //var ptr *int = 12 //doesnt work but works in C
 var ptr *int
 value := 13
 fmt.Println("Pointer default value ", ptr)
 ptr = &value
 fmt.Println("Pointer default value ", ptr)
 fmt.Println("Pointer data value ", *ptr)
 *ptr = *ptr+2
 fmt.Println("Pointer new data value ", *ptr)
 
}