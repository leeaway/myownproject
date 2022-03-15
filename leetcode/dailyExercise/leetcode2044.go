package dailyExercise

func countMaxOrSubsets(nums []int) int {
	/*
		二进制枚举解法
		maxIndex := 1<<len(nums)
		resMap := make(map[int]int)
		max := 0
		for i:=1;i<maxIndex;i++ {
			curRes := 0
			for j:=0;j<len(nums);j++ {
				if ((i>>j)&1) > 0 {
					curRes |= nums[j]
				}
			}
			resMap[curRes]++
			if max < resMap[curRes] {
				max = resMap[curRes]
			}
		}
		return max
	*/

	//回溯做法
	max := 0
	res := 0
	var dfs func(int, int)
	dfs = func(curPos int, curRes int) {
		if curPos == len(nums) {
			if max < curRes {
				max = curRes
				res = 1
			} else if max == curRes {
				res++
			}
			return
		}
		dfs(curPos+1, curRes|nums[curPos])
		dfs(curPos+1, curRes)
	}
	dfs(0, 0)
	return res
}
