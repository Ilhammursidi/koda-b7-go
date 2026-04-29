package goroutine

import (
	"fmt"
	"sync"
)

func Sender(board chan string, wg *sync.WaitGroup, pengirim, pesan string) { // pengirim
	defer wg.Done()
	board <- fmt.Sprintf("%s %s", pengirim, pesan)
}

func WhiteBoard(board chan string) { // menampilkan
	for msg := range board {
		fmt.Println(msg)
	}
}
