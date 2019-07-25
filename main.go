package main

import (
	"LearningGolang/part1"
)

func main()  {
	//call_part1()
	//part1.PrintAddress()
	//part1.TestAddress()
	//part1.Test13()
	//part1.TestString()
	//part1.TestObjectCompare()
	//part1.TestIota()
	//part1.TestSwitch()
	//part1.TestStringIndex()
	//part1.TestConcatByteArrayToString()
	//part1.TestBase58()
	//part1.GetUint32()
	//part1.TestSplit()
	//part1.TestSize()
	//part1.TestByteArr()
	//part1.TestStringCopy()
	//part1.TestHeap()
	//fmt.Println(part1.GetRandomCommonName())
	//fmt.Println(part1.GetRandomCommonName())
	//fmt.Println(part1.GetRandomCommonName())
	//fmt.Println(part1.GetRandomCommonName())
	part1.TestErrorf()
	//part1.TestConfig()
	//part1.TestStringRange()
	//part2.TestSlice()


}

func call_part1() {
	a,b  := make([]int, 100), make([]int ,100)
	//打印2^256 ,并且输出数的位数, 高精度加法
	part1.PowerBy2(256, a,b)
}






