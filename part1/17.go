package part1

import (
	"fmt"
	"github.com/xtario/xtar/common/errors"
	"math/rand"
	"time"
	"unsafe"

	"strconv"
	"strings"
)

type ConfigInfo struct {
	maxNamesCountLimit uint32
	stakeNamePriceList [16]uint64
}

var config *ConfigInfo

func GetConfig() *ConfigInfo {
	var err errors.Error

	if config == nil {
		config, err = NewConfig()
		if err != errors.ErrOK {
			config = &ConfigInfo{}
		}
	}
	return config
}

func NewConfig() (*ConfigInfo, errors.Error) {

	var config = new(ConfigInfo)
	namePriceStr := "0,10000,5000,3000,2000,1000,800,600,400,300,200,100,80,60,50,40,30"

	if len(namePriceStr) == 0 {
		return nil, errors.ErrCtrExecute.SetMsg(fmt.Sprintf("get global param stakeNamePriceList error:%s", "err"))
	}

	arrStr := strings.Split(namePriceStr, ",")

	for i, strPrice := range arrStr {

		if i > len(config.stakeNamePriceList) {
			break
		}

		var err error
		config.stakeNamePriceList[i], err = strconv.ParseUint(strPrice, 10, 64)
		if err != nil {
			return nil, errors.ErrCtrExecute.SetMsg(fmt.Sprintf("invalid global param stakeNamePriceList error:%s", err))
		}
	}

	return config, errors.ErrOK
}

func TestConfig() {
	GetConfig().maxNamesCountLimit++
	fmt.Println(GetConfig().stakeNamePriceList)
}

func TestSplit() {
	arr := strings.Split("", ".")
	fmt.Println(len(arr))
}

func TestSize() {
	str := "1234567890"
	fmt.Println(unsafe.Sizeof(str))
}

func TestByteArr() {
	str := "12345"
	size := len(str) + 1
	key := make([]byte, size)
	copy(key[1:], str[:])
	fmt.Println(key)
}

type Obj struct {
	str string
}

func TestStringCopy() {
	obj := Obj{}
	fmt.Println(obj.str == "")
}

func GetRandomCommonName() string {
	str := ""
	for i := 0; i < 12; i++ {
		str = str + string(rune(97+rand.Intn(26)))
	}
	return str
}

func TestMap() {
	a := make(map[string]int)

	fmt.Println(a["hyz"])
}

func TestSliceMake() {

	a := make([]int, 10)
	fmt.Println(a)

	b := make([]*Obj, 10)
	fmt.Println(len(b))
}

func TestUint() {
	var a, b uint32
	const c = 24 * 3600

	a = 1560756725
	b = 1560753446
	fmt.Println(a-b)
	fmt.Println(c)
	fmt.Println(a-b <= c)
}

func TestStringBool() {
	a:=true
	fmt.Println(strconv.FormatBool(a))
}

func TestPrint() {
	//a:=5
	fmt.Printf("%%\n")
}

func TestSlice() {
	a:=[]int{1}
	a=a[0:0]
	fmt.Println(a)
}

func TestRandomStr(){
	for i:=1; i<100; i++ {
		fmt.Println(getRandomCommonName())
	}
}

func getRandomCommonName() string {
	charString := "123456789abcdefghijklmnopqrstuvwxyz"
	str := ""

	for i := 0; i < 12; i++ {
		str = str + string(charString[rand.Intn(26)])
	}
	return str
}

func TestMapOrder()  {

	m := make(map[string]string)
	hash := make(map[string]int)

	m["hyz"] = "h"
	m["123"] = "456"
	m["345"] = "456"
	m["6"] = "456"
	m["7"] = "456"
	m["8"] = "456"
	m["8"] = "456"
	m["11"] = "456"
	m["11111"] = "456"
	m["ethan"] = "456"
	m["dsaf"] = "456"
	m["etgfhhan"] = "456"
	m["dfgsgsfd"] = "456"
	m["jlan"] = "456"
	m["zzz"] = "456"



	for i :=0; i<1000; i++ {
		for k,_ :=range m {
			hash[k] ++
			break
		}
	}

	for k,_ :=range m{
		fmt.Printf("%d, %s, \n", hash[k], k)
	}
}


func TestMapPointer() {
	m := make(map[string]*Obj)

 	fmt.Println(m["123"])

}

func TestUnixTimestamp() {
	//var bidTime uint32 = 1561105281
	fmt.Println(time.Unix(time.Now().Unix(),1000).Format("2006-01-02 15:04:05"))

}

func TestSep(){
	a:=10
	b:=3
	c:=a/b
	d:=float64(a)/float64(b)
	fmt.Println(a/b, c, d)
	e:=uint64(d*10)
	fmt.Println(e)
}