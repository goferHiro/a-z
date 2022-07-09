package main


import (
  "fmt"
  "io"
  "os"
  "ioutil"
)


func main(){
  fmt.Println("W")
  content := "Simple Haappy"
  file,err := os.Create("simple.txt")
  CheckNilError(err)
  length,err := io.WriteString(file, content)
  CheckNilError(err)
  fmt.Println("length is ",length)
}

func readFile(filename string){
  databyte,err:=ioutil.ReadFile(filename)
  CheckNilError(err)
  fmt.Println(string(databyte))
}

func CheckNilError(err error){
  if (err!=nil){
    panic(err)
  }
}

