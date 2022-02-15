package leetcode

func luckyNumbers(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	var rowMin, colMax []int
	for i := 0; i < m; i++ {
		rowMin = append(rowMin, getMin(matrix[i]))
	}
	for j := 0; j < n; j++ {
		curMax := 0
		for i := 0; i < m; i++ {
			if matrix[i][j] > curMax {
				curMax = matrix[i][j]
			}
		}
		colMax = append(colMax, curMax)
	}
	var res []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rowMin[i] == colMax[j] {
				res = append(res, rowMin[i])
			}
		}
	}
	return res
}

func getMin(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	return min
}
