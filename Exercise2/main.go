package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var dirSize int64 = 0

func readSize(path string, file os.FileInfo, err error) error {
	if !file.IsDir() {
		dirSize += file.Size()
	}
	return nil
}

func DirSizeMB(path string) float64 {
	dirSize = 0
	filepath.Walk(path, readSize)
	sizeMB := float64(dirSize) / 1024.0 / 1024.0
	return sizeMB
}

func main() {
	arg := os.Args[1]

	result := DirSizeMB(arg)
	//fmt.Println(result)
	var intResult int = int(result)
	fmt.Println(intResult)
}
