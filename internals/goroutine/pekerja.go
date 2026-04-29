package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func BangunPagi() {
	var wg sync.WaitGroup
	// var chn chan string
	// chn = make(chan string)

	// pekerja := 1
	// for range pekerja {
	wg.Add(1)
	go mandi(&wg)
	// wg.Wait()

	wg.Add(1)
	go sarapan(&wg)
	// wg.Wait()

	wg.Add(1)
	go rapikanTempatTidur(&wg)
	// wg.Wait()

	wg.Add(1)
	go buatKopi(&wg)
	// }
	// for i := range 5 {
	// 	do := fmt.Sprintf("Melakukan %d", i)
	// 	chn <- do
	// 	fmt
	// }

	// close(chn)
	wg.Wait()
	fmt.Println("berangkat kerja")

}

func mandi(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("Mulai mandi \n")
	time.Sleep(1 * time.Second)
	fmt.Printf("Selesai mandi \n")
}

func buatKopi(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Mulai buat kopi")
	time.Sleep(1 * time.Second)
	fmt.Println("Selesai buat kopi")
}

func sarapan(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Mulai membuat sarapan")
	time.Sleep(1 * time.Second)
	fmt.Println("Selesai membuat sarapan")
}

func rapikanTempatTidur(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Mulai rapikan tempat tidur")
	time.Sleep(1 * time.Second)
	fmt.Println("Selesai rapikan tempat tidur")
}
