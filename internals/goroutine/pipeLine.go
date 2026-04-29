package goroutine

func GetNumber(n int) <-chan int {
	numChan := make(chan int)
	go func() {
		for i := 1; i < n+1; i++ {
			numChan <- i
		}
		close(numChan)
	}()
	return numChan
}

func Genap(in <-chan int) <-chan int {
	numChan := make(chan int)
	go func() {
		for num := range in {
			if num%2 == 0 {
				numChan <- num
			}
		}
		close(numChan)
	}()
	return numChan
}
func Kuadrat(in <-chan int) <-chan int {
	numChan := make(chan int)
	go func() {
		for num := range in {
			numChan <- num * num
		}

		close(numChan)
	}()
	return numChan
}
