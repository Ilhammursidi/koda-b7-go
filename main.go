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

	// minitask 2
	var n int
	fmt.Print("Input angka : ")
	fmt.Scan(&n)
	triangle(n)

	// minitask 3
	slice := []int{50, 75, 66, 20, 32, 90}
	var newSlice []int = slice[:3]
	var endSlice []int = slice[3:]

	var allSlice []int = make([]int, 0, 7)
	allSlice = append(allSlice, newSlice...)
	allSlice = append(allSlice, 88)
	allSlice = append(allSlice, endSlice...)
	fmt.Println(slice)
	fmt.Println(newSlice)
	fmt.Println(endSlice)
	fmt.Println(allSlice)

	// minitask 4
	ilham := user{
		// id:        1,
		name:      "Ilham Mursidi",
		image:     "",
		email:     "imursidi37@gmail.com",
		umur:      26,
		phone:     "082189971471",
		isMarried: false,
		education: education{name: "SMKN 1 Minasatene", major: "Teknik Komputer dan Jaringan"},
	}

	fmt.Println(ilham.education)
	fmt.Println(ilham)
}

type user struct {
	// id        int
	name      string
	image     string
	email     string
	umur      int
	phone     string
	isMarried bool
	education education
}

type education struct {
	name  string
	major string
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
