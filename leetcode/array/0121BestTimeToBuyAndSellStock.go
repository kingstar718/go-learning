package main

import "fmt"

/**
给定一个数组，第i个元素是第i天股票的价格，最多只允许完成一次交易

可以用dp，也可以用单调栈
*/

func maxProfit(princes []int) int {
	if len(princes) < 1 {
		return 0
	}
	min, max := princes[0], 0
	for i := 0; i < len(princes); i++ {
		// 如果当前值-最小值比最大值大，赋予其最大值
		if princes[i]-min > max {
			max = princes[i] - min
		}
		// 如果当前值比最小值小，将最小值赋予
		if princes[i] < min {
			min = princes[i]
		}
	}
	return max
}

func maxProfit0121() {
	p := []int{7, 1, 5, 3, 6, 4}
	p2 := []int{7, 6, 4, 3, 2, 1}
	fmt.Println(maxProfit(p))
	fmt.Println(maxProfit(p2))
}
