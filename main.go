package main

import (
	"LearningGolang/part1"
	"LearningGolang/part2"
)

func main()  {
	//call_part1()
	//part1.PrintAddress()
	//part1.TestAddress()

	part2.TestSlice()
}


func call_part1() {
	a,b  := make([]int, 100), make([]int ,100)
	//打印2^256 ,并且输出数的位数, 高精度加法
	part1.PowerBy2(256, a,b)
}






