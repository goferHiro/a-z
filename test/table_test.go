package test

import "testing"

type tableTest struct{
   id int
   input string
   output string
}


func table_func (s string) (o string){
  return s
}


func TestTable(t *testing.T){
  inputs := []tableTest{
    {len(inputs),"abc","abc"},
    {len(inputs),"r","x"},
  }
  for i, input := range(inputs){
    if actualOut:= table_func(input.input);actualOut!=input.output{
      t.Errorf("testcase-id %v failed",input.id)
    }
  }
}