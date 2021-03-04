package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func fact2(n int) int {
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	return sum
}

func main() {
	fmt.Println(fact(7))
	fmt.Println(fact2(7))
}
