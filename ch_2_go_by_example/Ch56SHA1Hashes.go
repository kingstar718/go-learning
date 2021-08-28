package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	// 这将作为字节切片获得最终的哈希结果。参数 toSum可用于附加到现有的字节切片：通常不需要它。
	bs := h.Sum(nil)

	fmt.Println(s)
	// SHA1 值通常以十六进制打印，例如在 git commits 中。
	// 使用%x格式动词将哈希结果转换为十六进制字符串。
	fmt.Printf("%x\n", bs)
}
