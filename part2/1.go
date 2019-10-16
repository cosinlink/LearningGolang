package part2

import (
	"fmt"
	"runtime"
)

func TestSlice() {
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[1:5]

	fmt.Println(slice)
	fmt.Println(newSlice)
}

func PrintOs() {
	fmt.Println(runtime.GOOS)
	//fmt.Println(runtime.GOARCH)

}
