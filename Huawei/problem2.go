package Huawei

/**
 * @author 2416144794@qq.com
 * @date 2023/2/9 21:59
 */

/*
	ropes: 所有绳子
	target: 需要裁出多少根绳子
	@return: 返回可以切出target根绳子的最长长度
*/
func solveProblem2(ropes []int, target int) int {
	maxRope, sumLen := 0, 0
	for _, rope := range ropes {
		sumLen += rope
		if rope > maxRope {
			maxRope = rope
		}
	}

	maxLen := sumLen / target
	if maxLen == 0 {
		return 0
	}

	//这里用二分搜索
	l, r := 0, maxLen

	for l <= r {
		mid := l + (r-l)/2
		if checkCanCut(ropes, target, mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}

func checkCanCut(ropes []int, target, eachLen int) bool {
	res := 0
	for _, rope := range ropes {
		res += rope / eachLen
		if res >= target {
			return true
		}
	}
	return false
}
