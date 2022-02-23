package leetcode

func NumberOfGoodSubsets(nums []int) int {
	banNum := []int{4, 8, 9, 12, 16, 18, 20, 24, 25, 27, 28}
	primeNum := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	banMap, primeMap := make(map[int]bool), make(map[int]bool)
	for _, num := range banNum {
		banMap[num] = true
	}
	for _, num := range primeNum {
		primeMap[num] = true
	}

	numToCountMap := make(map[int]int)
	var validNums []int
	for _, num := range nums {
		if banMap[num] {
			continue
		}
		validNums = append(validNums, num)
		numToCountMap[num]++
	}
	res := 0
	for i := 1; i < (1 << len(validNums)); i++ {
		var curNumList []int
		for index, num := range validNums {
			if ((i >> index) & 1) > 0 {
				curNumList = append(curNumList, num)
			}
		}
		if isValid(curNumList, primeMap) {
			res++
		}
	}
	return res
}

func isValid(nums []int, primeMap map[int]bool) bool {
	if len(nums) == 1 && nums[0] == 1 {
		return false
	}
	existMap := make(map[int]bool)
	for _, num := range nums {
		for _, n := range resolvePrime(num, primeMap) {
			if existMap[n] {
				return false
			}
			existMap[n] = true
		}
	}
	return true
}

func resolvePrime(a int, primeMap map[int]bool) []int {
	start := 2
	var res []int
	//a是质数
	if primeMap[a] {
		return []int{a}
	}
	for start < a {
		if a%start == 0 {
			a = a / start
			res = append(res, start)
			continue
		}
		start++
	}
	return res
}
