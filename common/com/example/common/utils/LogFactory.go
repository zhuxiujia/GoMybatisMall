package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func NewOutPut(logpath string) io.Writer {
	logFile := createFileWithDir(logpath)
	return &MyWrite{file: logFile}
}

type MyWrite struct {
	file *os.File
}

func (it *MyWrite) Write(p []byte) (n int, err error) {
	//TODO 实际项目中 其实不应该输出到控制台，这里可以注释 fmt.Println()
	fmt.Println(string(p))
	return it.file.Write(p)
}

func createFileWithDir(name string) *os.File {
	var path = name
	if !strings.Contains(path, ".log") {
		panic("日志文件必须以.log结尾")
	}
	path = path[0:strings.LastIndex(path, "/")]
	os.MkdirAll(path, os.ModePerm)
	file, logErr := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if logErr != nil {
		fmt.Println("Fail to find", name, logErr)
		os.Exit(1)
	}
	return file
}
