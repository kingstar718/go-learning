package main

import "fmt"

/**
给定一个非负整数，生成杨辉三角的前numRows行
每个数是它左上方和右上方之和
*/

func generate(numRows int) [][]int {
	var result [][]int
	for i := 0; i < numRows; i++ {
		var row []int
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else if i > 1 {
				row = append(row, result[i-1][j-1]+result[i-1][j])
			}
		}
		result = append(result, row)
	}
	return result
}

func generate0118() {
	fmt.Println(generate(5))
}
