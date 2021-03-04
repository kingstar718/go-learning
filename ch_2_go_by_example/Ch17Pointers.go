package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

/**
& 取地址运算符 &i表示取变量i的地址
* 间接寻址符 是&运算符的补充，返回操作数所制定地址的变量的值 即*(&i)还是i的值
*/

func main() {
	i := 0
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
