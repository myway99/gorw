package main

import (
	"bufio"
	"fmt"
	"os"
)

//命令行：exe main.go
func main() {

	//检查输入的参数os.Args
	//判定os.Args切片大小，若小于2则返回
	if len(os.Args) < 2 {
		return
	}

	//取os.Args切片的第二个（下标为1）作为文件名
	filename := os.Args[1]

	//打开文件
	file, err := os.Open(filename)
	//若打开文件出错，则返回
	if err != nil {
		fmt.Println(err)
		return
	}
	//延时到main函数结束时关闭文件，回收文件句柄
	defer file.Close()

	//通过bufio读取文件
	reader := bufio.NewReader(file)

	var line int

	//循环读取文件的每一行
	for {

		//对超宽行的文件兼容
		_, isPrefix, err := reader.ReadLine()

		if err != nil {
			break
		}

		//统计行数
		if !isPrefix {
			line++
		}

	}

	//打印统计结果
	fmt.Println(line)

}
