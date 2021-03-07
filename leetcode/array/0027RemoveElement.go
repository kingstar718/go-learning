package main

import "fmt"

/**
只是将要删除的元素移动到数组后面的空间内，然后返回数组实际的剩余空间
*/
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return j
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
}
