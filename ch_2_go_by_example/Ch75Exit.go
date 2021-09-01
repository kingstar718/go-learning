package main

import (
	"fmt"
	"os"
)

func main() {
	//defer不会在使用时运行os.Exit，因此fmt.Println永远不会被调用。
	defer fmt.Println("!")

	os.Exit(3)
}
