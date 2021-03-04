package main

import "fmt"

/**
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
*/

/**
给出一个数组，代表一个十进制数，数组的 0 下标是十进制数的高位。
要求计算这个十进制数加一以后的结果。
简单的模拟题。从数组尾部开始往前扫，逐位进位即可。
最高位如果还有进位需要在数组里面第 0 位再插入一个 1 。
*/

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+carry > 9 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i] += carry
			carry = 0
		}
	}
	// 如果加到最后，首位为0，代表为999之类的数字，以上得到的000->1000
	if digits[0] == 0 && carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 3, 2, 1}
	nums3 := []int{9, 9, 9}
	fmt.Println(plusOne(nums1))
	fmt.Println(plusOne(nums2))
	fmt.Println(plusOne(nums3))
}
