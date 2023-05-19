package main

import (
	"fmt"
	snbt "github.com/nas-cod/mc-snbt"
	"os"
)

// ReadFileToBytes 打开文件并读取Bytes
func ReadFileToBytes(filepath string) ([]byte, error) {
	file, err := os.ReadFile(filepath)
	return file, err
}

func main() {
	bytes, err := ReadFileToBytes("test_data/creative.snbt")
	if err != nil {
		return
	}

	var m1 = map[string]interface{}{}

	err = snbt.Decode(bytes, &m1)
	if err != nil {
		return
	}

	fmt.Println(m1)
}
