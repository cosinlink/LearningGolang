package part1
var aa,bb,cc,dd,ee,ff int32

type Person struct {
	age int32
}

func PrintAddress() *Person{
	var arr [2]int32
	var a,b,c,d,e,f int32
	person := &Person{ 11 }

	println(&a, &b, &c, &d, &e, &f)
	println(&arr, &arr[0], &arr[1])
	println(person)

	println(&aa, &bb, &cc, &dd, &ee, &ff)
	return person
}

func TestAddress() {
	p := PrintAddress()
	println(p)
}
