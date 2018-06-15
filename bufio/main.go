package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//读buffer
	strReader := strings.NewReader("hello world")

	bufReader := bufio.NewReader(strReader)
	//提取
	data, _ := bufReader.Peek(5)
	fmt.Println(data, bufReader.Buffered())
	//抓取
	str, _ := bufReader.ReadString(' ')
	fmt.Println(str, bufReader.Buffered())

	//写buffer
	w := bufio.NewWriter(os.Stdout)

	fmt.Fprint(w, "Hello ")
	fmt.Fprint(w, "world!")
	//提交
	w.Flush()

}
