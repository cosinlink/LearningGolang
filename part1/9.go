package part1

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const limit  =  1000
// this is github.com
/*
	压位高精度加法, limit为单个数的最大上限
	a + b = c
	a[0] 表示切片的长度, a[1] - a[a[0]] 分别为从低位到高位,每个数不超过limit
*/
func Add_big(a,b,c []int ) {
	length :=a[0]
	//fmt.Println(length)

	i, tmp, delta:=1, 0, 0

	for i <= length+1 {
		if a[i]>=limit || b[i] >=limit {
			log.Fatal("number can not more than limit_" + strconv.Itoa( limit ) )
		}
		tmp = a[i] + b[i] + delta
		delta = 0
		if (tmp >= limit) {
			delta = 1
			tmp -= limit
		}

		c[i] = tmp
		i++
	}
	if c[length+1]>0 {
		c[0] = length +1
	}

}


/*
	求2^num,通过压位高精度求解,并且打印出来,有很多细节需要注意
	比如打印的时候 前面要加0
*/
func PowerBy2(num int, a,b []int)  {
	a[0] = 1
	a[1] = 1


	for i:=1; i<=num; i++ {
		copy(b, a)
		Add_big( a, b, a )
	}

	fmt.Printf("2^%d:  ", num )
	printX(a)
}


func printX(a []int) {
	result := ""
	zeroLen:=0

	for i := 1; i < a[0]; i++ {
		zeroLen = 3 - len(strconv.Itoa(a[i]))
		result = strings.Repeat("0", zeroLen) + strconv.Itoa(a[i]) + result
	}

	result = strconv.Itoa( a[a[0]] ) + result
	fmt.Println(result + "   length=" , len(result))
}

