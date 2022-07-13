package generators


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

