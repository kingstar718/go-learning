package main

import "fmt"

/**
给定一个数组，第i个元素是第i天股票的价格，允许多次交易且必须买之前卖出股票
最大收益来源，必然是每次跌了就买入，涨到顶峰的时候就抛出。只要有涨峰就开始计算赚的钱，
连续涨可以用两两相减累加来计算，两两相减累加，相当于涨到波峰的最大值减去谷底的值。这一点看通以后，题目非常简单。
*/

func maxProfit2(princes []int) int {
	profit := 0
	for i := 0; i < len(princes)-1; i++ {
		if princes[i+1] > princes[i] {
			profit += princes[i+1] - princes[i]
		}
	}
	return profit
}

func main() {
	p := []int{7, 1, 5, 3, 6, 4}
	p2 := []int{7, 6, 4, 3, 2, 1}
	fmt.Println(maxProfit2(p))
	fmt.Println(maxProfit2(p2))
}
