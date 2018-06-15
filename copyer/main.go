package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

//判断文件是否存在
func fileExists(filename string) bool {

	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//拷贝文件
func copyFile(src, dst string) (w int64, err error) {

	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

//拷贝文件的行为
func copyFileAction(src, dst string, showProgress, force bool) {

	if !force {
		if fileExists(dst) {

			fmt.Printf("%s exists override? y/n\n", dst)
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()

			if strings.TrimSpace(string(data)) != "y" {
				return

			}

		}

	}

	copyFile(src, dst)

	if showProgress {

		fmt.Printf("'%s'->'%s'\n", src, dst)

	}

}

func main() {

	var showProgress, force bool

	//定义命令行参数
	flag.BoolVar(&force, "f", false, "force copy when existing")
	flag.BoolVar(&showProgress, "v", false, "explain what is being done")

	flag.Parse()

	//非法命令行数检测：判断命令行参数个数
	if flag.NArg() < 2 {
		flag.Usage()
		return
	}

	//调用拷贝文件的行为
	copyFileAction(flag.Arg(0), flag.Arg(1), showProgress, force)

}
