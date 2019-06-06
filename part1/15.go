package part1

import (
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	com "github.com/BXA/bxa/common"
	"strconv"
)

const AddressSize = 20

//ToBase58 return address in base58 code
func (a Address) ToBase58() string {
	d := append([]byte{23}, a[:]...)
	tmp1 := com.Hash(d)
	tmp2 := com.Hash(tmp1[:])
	d = append(d, tmp2[0:4]...) //check sum
	e, err := base58.BitcoinEncoding.Encode([]byte(new(big.Int).SetBytes(d).String()))
	if err != nil {
		panic(err)
	}
	return string(e)
}

func TestBase58()  {
	add:= Address{}
	fmt.Println(add.ToBase58())
}

func TestLenNil() {
	str := ""

	fmt.Println(len(str))
}

func GetUint32() (uint32, error) {

	u, err := strconv.ParseUint(string(""), 10, 32)
	if err != nil {
		return 0, err
	}
	fmt.Println(uint32(u))
	return uint32(u), nil
}