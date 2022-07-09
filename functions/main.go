package main


func adder(val1 int ,val2 int) int{
  return val1+val2
}

func variadicAdder(vals ...int) int{
  total:=0
  for _, val:=range total{
    total+=val
  }
  return total
}

func main(){

  _:= adder(1,2)
  
}