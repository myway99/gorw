package main

import (
	"fmt"
	"os"
)

//定义一个结构体
type Data struct {
}

//声明一个String的成员函数，返回值一个字符串
//将OO里的this直接命名成self
func (self Data) String() string {

	return "data"

}

func main() {

	//格式化输出（默认不带回车）
	fmt.Printf("hello world\n")

	//没格式化功能（自带回车），可带较多参数,并逐个打印出来，较更好用
	fmt.Println("myway", 2, "str")

	//%v可自动识别类型
	fmt.Printf("num %v\n", "h")

	//把格式化好的字符串输出成返回值
	str := fmt.Sprintf("float %f", 3.14159)
	//简单打印
	fmt.Print(str)

	//指定IO writer输出到指定的接口
	fmt.Fprintln(os.Stdout, "A\n")

	//调用函数得到结构体并打印
	fmt.Printf("%v\n", Data{})

}
