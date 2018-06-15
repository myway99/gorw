//main
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	s := "hello world"
	//Contains()函数判断是否包含
	fmt.Println(strings.Contains(s, "hello"), strings.Contains(s, "?"))
	//其实是查看索引,base0
	fmt.Println(strings.Index(s, "o"))

	ss := "1#2#345"

	//Split切割字符串
	splitedStr := strings.Split(ss, "#")
	fmt.Println(splitedStr)

	//Join合并字符串
	fmt.Println(strings.Join(splitedStr, "#"))

	//检查前缀HasPrefix和后缀HasSuffix
	fmt.Println(strings.HasPrefix(s, "he"), strings.HasSuffix(s, "ld"))

	//字符串转换
	//基本数值转换
	//i到a
	fmt.Println(strconv.Itoa(10))
	//a到i
	fmt.Println(strconv.Atoi("711"))

	//Parse解析
	fmt.Println(strconv.ParseBool("false"))

	fmt.Println(strconv.ParseFloat("3.14", 32))
	fmt.Println(strconv.ParseFloat("3.14", 64))

	//Format格式化
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(123, 10))
	fmt.Println(strconv.FormatInt(123, 2))
	fmt.Println(strconv.FormatInt(123, 4))
	fmt.Println(strconv.FormatInt(20, 16))

}
