package main

import "fmt"

/**
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。
*/
/**
经典的二分搜索的变种题，在有序数组中找到最后一个比target小的元素
*/

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] >= target {
			high = mid - 1
		} else {
			// 结束条件：mid等于数组长度时或左边的值大于等于target的时候
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			low = mid + 1
		}
	}
	return 0
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5)) //2
	fmt.Println(searchInsert(nums, 2)) //1
	fmt.Println(searchInsert(nums, 7)) //4
	fmt.Println(searchInsert(nums, 0)) //0
}
