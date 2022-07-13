package test

import (
	"math/rand"
	"testing"
)

type tableTest struct{
	id int
	input string
	output string
}


func table_func (s string) (o string){
	return s
}

func generateRandomNumbers () (ch chan int) {
	ch = make (chan int)

	go func(){
		for {
			ch <- rand.Int()
		}
	}()

	return
}

func next(ch chan int) (int){
	return <-ch
}

func TestTable(t *testing.T){
	numbers := generateRandomNumbers()
	inputs := []tableTest{
		{next(numbers),"abc","abc"},
		{next(numbers),"r","x"},
	}
	close(numbers)
	for _, input := range(inputs){
		if actualOut:= table_func(input.input);actualOut!=input.output{
			t.Errorf("testcase-id %v failed",input.id)
		}
	}
}