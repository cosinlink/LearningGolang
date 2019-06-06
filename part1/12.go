package part1

import (
	"fmt"
	"reflect"
)

func Test1()  {

	param := []int{10,29}
	fmt.Println(  reflect.TypeOf(param)  )

	test2( 1,2,3 )
	test3( param )
	test5( param... )
	//test5( param )
}

func test2( param ...interface{} )  {
	fmt.Println(param)

	fmt.Println(  reflect.TypeOf(param)  )

	test3(param...)
}

func test3(param ...interface{})  {
	fmt.Println(param)
}


func test4(param []interface{})  {
	fmt.Println(param)
}

func test5(param ...int)  {
	fmt.Println(param)
}