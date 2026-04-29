package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func CoffeShop() {
	var wg sync.WaitGroup
	// var chn chan string
	var chn chan string
	// chn = make(chan string)
	chn = make(chan string, 3) //buffer
	// coffees := []string{"kopi1", "kopi2", "kopi3", "kopi4"}

	baristaCount := 3
	for range baristaCount {

		// wg.Go(func() {
		// 	Barista(chn, &wg)
		// })
		wg.Add(1)
		go Barista(chn, &wg)
	}

	for i := range 12 {
		order := fmt.Sprintf("kopi %d", i)
		chn <- order
		fmt.Printf("Order masuk %s\n", order) // pertama
	}
	close(chn)
	wg.Wait()
	fmt.Println("Toko Tutup")

}

func Barista(coffeChn chan string, wg *sync.WaitGroup) {
	// defer wg.Done()
	defer func() {
		fmt.Println("Barista Pulang")
		wg.Done()
	}()
	for coffee := range coffeChn {
		time.Sleep(1 * time.Second)
		fmt.Printf("Start making coffe %s\n", coffee)
		time.Sleep(1 * time.Second)
		fmt.Printf("Finish making coffe %s\n", coffee)
	}

}
