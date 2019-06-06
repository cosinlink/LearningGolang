package part1

import (
	"fmt"
	"strings"
)


const (
	Init         = "init"
	RegisterName = "registerName"
	QueryAddress = "queryAddress"
)

const (
	x = iota
	y = iota
	z = iota
)

const (
	xx = iota
	yy = iota
	zz = iota
)


func TestString(){

	a := "abcdef"
	b := make([]byte,10,10)
	copy(b, a[:])
	fmt.Println(b)

}

var EmptyAddress = Address{}
type Address [20]byte


func TestObjectCompare() {
	a:= EmptyAddress
	a[0] =1
	b:= Address{}
	b[0] =1
	fmt.Println( &a == &b )
}

func TestStringRange()  {

	arr := []string {"xueyuozieos3",  "enbank.io", "moonriverpub", "gm2tkmbyguge", "cjpxmh521314", "!dfsajhfkj11", "1", "!" }
	for i, str := range arr{

		fmt.Println(i,str, ValidateCommonName(str))
	}
}

func ValidateCommonName(nameContent string) bool {
	arr := []rune(nameContent)

	if len(arr) != 12 {
		return false
	}

	for _, ch := range arr {
		if (ch>='a' && ch<='z') || (ch>='1' && ch<='5'){
			continue
		} else{
			return false
		}
	}
	return true
}

func TestIota() {
	fmt.Println(x,y,z,xx,yy,zz)
}

func TestSwitch() {
	x:= 0
	switch x {
	case 0:
		fmt.Println("a")
	case 1:
		fmt.Println("b")
	default:
		fmt.Println("xxx")

	}
}

func TestStringIndex()  {
	fmt.Println(strings.Index("widuu", "i")) //1
	fmt.Println(strings.Index("wuiduu", "u")) //3
	fmt.Println(strings.IndexAny("wuiduu", "u")) //3
}

func TestConcatByteArrayToString() {
	a:= []byte{97,98,99,100,49,49,48}
	b:=make([]byte, 0)
	fmt.Println(string(a))
	fmt.Println(string(b) == "")


}

