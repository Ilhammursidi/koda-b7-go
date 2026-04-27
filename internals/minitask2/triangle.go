package minitask2

import "fmt"

func Triangle(n int) int {

	for i := 0; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	return n
}
