package main

import (
	//binary包
	"encoding/binary"
	"fmt"
	"os"
)

//声明结构体
type BitmapInfoHeader struct {
	Size           uint32
	Width          int32
	Height         int32
	Places         uint16
	BitCount       uint16
	Compression    uint32
	SizeImage      uint32
	XperlsPerMeter int32
	YperlsPerMeter int32
	ClsrUsed       uint32
	ClrImportant   uint32
}

func main() {

	//此处默认打开文件时没有出错
	file, _ := os.Open("splash.bmp")
	defer file.Close()

	//将一个WORD切成两个byte来读
	var headA, headB byte
	//binary包对Go二进制接口有很好的封装，会自动识别所输入的数据的类型并读取
	//字节序ByteOrder：
	//win和linux常用LittleEndian(更多文件格式）
	//mac常用BigEndian
	binary.Read(file, binary.LittleEndian, &headA)
	binary.Read(file, binary.LittleEndian, &headB)

	fmt.Printf("%c%c\n", headA, headB)

	//size
	var size uint32
	binary.Read(file, binary.LittleEndian, &size)

	//reserved
	var reservedA, reservedB uint16
	binary.Read(file, binary.LittleEndian, &reservedA)
	binary.Read(file, binary.LittleEndian, &reservedB)

	//offbits
	var offbits uint32
	binary.Read(file, binary.LittleEndian, &offbits)

	fmt.Println(headA, headB, size, reservedA, reservedB, offbits)

	//使用结构体来读取
	//先将infoHeader实例化
	infoHeader := new(BitmapInfoHeader)
	//调用binary从file里用LittleEndian读取，直接将infoHeader传入
	binary.Read(file, binary.LittleEndian, infoHeader)

	fmt.Println(infoHeader)

}
