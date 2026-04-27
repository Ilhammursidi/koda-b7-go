package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(luas(5))
	fmt.Println(keliling(5))
}

func luas(r float64) float64 {
	return math.Pi * r * r
}

func keliling(r float64) float64 {
	return 2 * math.Pi * r
}
