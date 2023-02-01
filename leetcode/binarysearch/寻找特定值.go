package binarysearch

/**
 * @author 2416144794@qq.com
 * @date 2023/2/1 17:50
 */

func binarySearchTarget(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

/*
	寻找等于target的左边界
*/
func binarySearchTargetLeftBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			r = mid - 1
		}
	}
	if l == len(nums) || nums[l] != target {
		return -1
	}
	return l
}

/*
寻找右边界
*/

func binarySearchTargetRightBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l-1 < 0 || nums[l-1] != target {
		return -1
	}
	return l - 1
}
