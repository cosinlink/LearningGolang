package part1

import (
	"fmt"
	"math"
	"strconv"
)

func PrintMaxUint() {
	fmt.Println("MaxUint8 ", strconv.FormatUint(uint64(math.MaxUint8), 10))
	fmt.Println("MaxUint16 ", strconv.FormatUint(uint64(math.MaxUint16), 10))
	fmt.Println("MaxUint32 ", strconv.FormatUint(uint64(math.MaxUint32), 10))
	fmt.Println("MaxUint64 ", strconv.FormatUint(math.MaxUint64, 10) )

	a:=true

	fmt.Println(!(a || true))
}
