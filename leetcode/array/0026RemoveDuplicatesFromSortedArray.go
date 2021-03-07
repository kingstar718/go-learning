package main

import "fmt"

/**
给定一个有序数组 nums，对数组中的元素进行去重，使得原数组中的每个元素只有一个。最后返回去重以后数组的长度值。
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}

/**
从nums的最后一个开始遍历, 然后与前一个进行对比.
如果相等, 则删除该位置的数.
不等, 则继续往前遍历.
*/
func removeDuplicates2(nums []int) int {
	for i := len(nums); i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}

func removeDuplicates3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 1
	preNum := nums[0]
	for _, v := range nums {
		if v != preNum {
			preNum = v
			nums[count] = v
			count += 1
		}
	}
	return count
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums3 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	//fmt.Println(removeDuplicates2(nums2))
	fmt.Println(removeDuplicates3(nums3))
}
