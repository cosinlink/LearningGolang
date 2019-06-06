package part1

import (
	"fmt"
	"reflect"
)

func Test13()  {
	a := []int {1,2,4}
	b := []int {1,2,3}

	fmt.Println(reflect.DeepEqual(a,b))

}
