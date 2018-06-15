package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//定义可提取节点里的属性值的函数
func getAttributeValue(attr []xml.Attr, name string) string {
	for _, a := range attr {

		if a.Name.Local == name {
			return a.Value
		}
	}

	return ""
}

func main() {

	//使用ioutil.ReadFile将文件读入内存，此时为byte数组
	content, err := ioutil.ReadFile("toml.xml")
	//进行xml解析,把文本流解析成token
	decoder := xml.NewDecoder(bytes.NewBuffer(content))

	var t xml.Token

	//配置状态机inContexts，并设置当前状态
	var inContexts bool

	//用t来接收不停的遍历得到的token
	//token返回的内容是每一个xml的节点
	//对xml的节点进行解析处理
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {

		switch token := t.(type) {

		case xml.StartElement:

			name := token.Name.Local
			//fmt.Println(name)

			if inContexts {

				if name == "context" {
					//找到的RegExpr后打印输出
					//fmt.Println(name)
					fmt.Println(getAttributeValue(token.Attr, "name"))
				}

			} else {
				if name == "contexts" {

					//先找到contexts后，修改状态机inContexts的状态
					inContexts = true
				}
			}

		//处理结尾
		case xml.EndElement:
			if inContexts {
				if token.Name.Local == "contexts" {

					//状态机复位
					inContexts = false
				}
			}

		}
	}

}
