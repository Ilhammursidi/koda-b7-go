package minitask6

import (
	"bufio"
	"fmt"
	"os"
)

func OpenFile(path string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("PANIC: ", path, r)
			fmt.Println("continue..")
		}
	}()

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("gagal menemukan file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic("gagal membaca file")
	}
}
