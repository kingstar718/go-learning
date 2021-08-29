package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	f, err := os.CreateTemp("", "sample")
	if err != nil {
		panic(err)
	}

	fmt.Println("Temp file name:", f.Name())
	// 完成后清理文件，明确地执行此操作是一种很好地做法
	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	if err != nil {
		panic(err)
	}
	// 写入很多临时文件时最好创建临时目录，返回位目录名称而不死打开文件
	dname, err := os.MkdirTemp("", "sampledir")
	if err != nil {
		panic(err)
	}
	fmt.Println("Temp dir name: ", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	if err != nil {
		panic(err)
	}
}
