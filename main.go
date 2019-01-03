package main

import "LearningGolang/part1"

func main()  {
	call_part1()
}

func call_part1() {
	a,b  := make([]int, 100), make([]int ,100)
	//打印2^1 -- 2^256 每个数,并且输出数的位数, 高精度
	part1.PowerBy2(256, a,b)
}


