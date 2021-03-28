package main

import "fmt"

/**
找出两个数之和等于 target 的两个数字，要求输出它们的下标。注意一个数字不能使用 2 次。下标从小到大输出。假定题目一定有一个解。
这一题比第 1 题 Two Sum 的问题还要简单，因为这里数组是有序的。可以直接用第一题的解法解决这道题。
*/

func twoSum167(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{-1, -1}
}

func main() {
	num := []int{2, 7, 11, 5}
	fmt.Println(twoSum167(num, 9))
}
