package part1

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const (
	GenesisTimestamp = uint32(1561910400) //2019-07-01 00:00:00
)

func calVoteWeight(now uint32, stakeWeight uint64) uint64 {
	precision := math.Pow10(4)
	weight := math.Floor(float64(now-GenesisTimestamp)/(7*24*3600*52)*precision) / precision
	return stakeWeight * uint64(math.Pow(2, float64(weight)))
}

func testCalToFile() {

}

func main() {
	fmt.Println("---------------")
	log.Println("------ log printl ----")
	func_log2file()
	func_log2fileAndStdout()
}

func func_log2file() {
	//创建日志文件
	f, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//完成后，延迟关闭
	defer f.Close()
	// 设置日志输出到文件
	log.SetOutput(f)
	// 写入日志内容
	log.Println("check to make sure it works")
}

func func_log2fileAndStdout() {
	//创建日志文件
	f, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//完成后，延迟关闭
	defer f.Close()
	// 设置日志输出到文件
	// 定义多个写入器
	writers := []io.Writer{
		f,
		os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	// 创建新的log对象
	logger := log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
	// 使用新的log对象，写入日志内容
	logger.Println("--> logger :  check to make sure it works")
}
