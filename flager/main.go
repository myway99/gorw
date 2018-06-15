package main

import (
	"flag"
	"fmt"
	"os"
)

func style1() {

	//格式化定义
	methodPtr := flag.String("method", "default", "method of sample")
	valuePtr := flag.Int("value", -1, "value of sample")

	//解析
	flag.Parse()

	fmt.Println(*methodPtr, *valuePtr)

}

func style2() {

	var method string
	var value int

	flag.StringVar(&method, "method", "default", "method of sample")
	flag.IntVar(&value, "value", -1, "value of sample")

	flag.Parse()

	fmt.Println(method, value)

}

func main() {

	//命令行解析
	//os.Args读取命令行参数
	fmt.Println(os.Args)

	style2()

}
