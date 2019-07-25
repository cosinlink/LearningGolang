package part1

import "fmt"

func TestErrorf() {
	fmt.Println("err: %+v",fmt.Errorf("12345"))
	fmt.Println("after error")

}