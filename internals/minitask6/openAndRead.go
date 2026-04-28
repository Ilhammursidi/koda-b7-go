package minitask6

import (
	"bufio"
	"fmt"
	"os"
)

func OpenFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic("gagal menemukan file")
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
