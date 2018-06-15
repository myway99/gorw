package main

import (
	"encoding/xml"
	"fmt"
)

//人物档案
type person struct {
	//Name string
	Name string `xml:"姓名,attr"`
	//Name string `xml:"name,attr"`
	//Age int
	Age int `xml:"年龄"`
	//Age int `xml:"Age"`
}

func main() {

	//结构体实例化
	p := person{Name: "davy", Age: 18}

	var data []byte
	var err error

	fmt.Println("将结构体解析为XML:")
	//Marshal将结构体p序列化解析为XML流data
	//MarshalIndent加前缀和缩进
	if data, err = xml.MarshalIndent(p, "", " "); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	p2 := new(person)

	fmt.Println("将XML生成结构体:")
	//Unmarshal将XML流data反序列化为结构体p2
	if err = xml.Unmarshal(data, p2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2)

}
