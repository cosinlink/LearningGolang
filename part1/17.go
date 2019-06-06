package part1

import (
	"fmt"
	"github.com/BXA/bxa/common/errors"
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
	namePriceStr :=  "0,10000,5000,3000,2000,1000,800,600,400,300,200,100,80,60,50,40,30"

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

func TestSplit()  {
	arr := strings.Split(".1kjlj",".")
	fmt.Println(len(arr))
}

func TestSize()  {
	str:= "1234567890"
	fmt.Println(unsafe.Sizeof(str))
}

func TestByteArr() {
		str := "12345"
		size := len(str)+1
		key:=make([]byte, size)
		copy(key[1:], str[:])
		fmt.Println(key)
}

type Obj struct {
	str string
}
func TestStringCopy() {
	obj:=Obj{}
	fmt.Println(obj.str == "")
}
