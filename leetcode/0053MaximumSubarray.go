package main

import "fmt"

/**
给定一个数组，找到一个具有最大和的连续子数组（最少包含一个元素）	，返回其最大和
*/

/**
这一题可以用 DP 求解也可以不用 DP。
题目要求输出数组中某个区间内数字之和最大的那个值。
dp[i] 表示 [0,i] 区间内各个子区间和的最大值，
状态转移方程是 dp[i] = nums[i] + dp[i-1] (dp[i-1] > 0)，
dp[i] = nums[i] (dp[i-1] ≤ 0)。
*/

//DP
func maxSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = max(res, dp[i])
	}
	return res

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y

}

func maxSubarray2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, res, p := nums[0], 0, 0
	for p < len(nums) {
		res += nums[p]
		if res > maxSum {
			maxSum = res
		}
		if res < 0 {
			res = 0
		}
		p++
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubarray(nums))  //6
	fmt.Println(maxSubarray2(nums)) //6
}
