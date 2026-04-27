package minitask3

import "fmt"

func AddNumber(n int) int {
	slice := []int{50, 75, 66, 20, 32, 90}
	var newSlice []int = slice[:3]
	var endSlice []int = slice[3:]

	var allSlice []int = make([]int, 0, 7)
	allSlice = append(allSlice, newSlice...)
	allSlice = append(allSlice, n)
	allSlice = append(allSlice, endSlice...)
	for _, value := range allSlice {
		fmt.Println(value)
	}
	return n
}
