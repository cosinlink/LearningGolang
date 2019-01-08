package part2

import "fmt"

func TestSlice() {
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[1:5]

	fmt.Println(slice)
	fmt.Println(newSlice)
}