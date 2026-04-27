package main

import (
	"fmt"
	"math"
)

func main() {
	// minitask 1
	fmt.Println(luas(5))
	fmt.Println(keliling(5))

	fmt.Println(luasDanKeliling(5))

	var n int
	fmt.Println("Input angka : ")
	fmt.Scan(&n)
	triangle(n)
}

func luas(r float64) float64 {
	return math.Pi * r * r
}

func keliling(r float64) float64 {
	return 2 * math.Pi * r
}

func luasDanKeliling(r float64) (luas float64, keliling float64) {
	luas = math.Pi * r * r
	keliling = 2 * math.Pi * r
	return luas, keliling
}

func triangle(n int) int {

	for i := 0; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	return n
}
