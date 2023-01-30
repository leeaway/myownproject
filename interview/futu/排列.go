package futu

/**
 * @author 2416144794@qq.com
 * @date 2023/1/16 19:27
 */

var res [][]int

func Solve(nums []int) [][]int {
	n := len(nums)
	used := make([]bool, n)
	solveHelper(nums, used, []int{})
	return res
}

func solveHelper(nums []int, used []bool, tmp []int) {
	n := len(nums)
	if len(tmp) == n {
		res = append(res, tmp)
		return
	}

	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		tmp = append(tmp, nums[i])
		solveHelper(nums, used, tmp)
		used[i] = false
		tmp = tmp[:len(tmp)-1]
	}

}
