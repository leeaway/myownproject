package dailyExercise

import "sort"

func countRectangles(rectangles [][]int, points [][]int) []int {
	yToXMap := make(map[int][]int)
	var ys []int
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i][0] < rectangles[j][0]
	})
	for _, rectangle := range rectangles {
		x, y := rectangle[0], rectangle[1]
		yToXMap[y] = append(yToXMap[y], x)
	}
	for y, _ := range yToXMap {
		ys = append(ys, y)
	}
	sort.Slice(ys, func(i, j int) bool {
		return ys[i] < ys[j]
	})
	res := make([]int, len(points))
	for i, point := range points {
		res[i] = cal(yToXMap, ys, point)
	}
	return res
}

func cal(ymap map[int][]int, ys []int, point []int) int {
	x, y := point[0], point[1]
	start := binarySearch(y, ys)
	sum := 0
	for i := start; i < len(ys); i++ {
		sortx := ymap[ys[i]]
		sx := binarySearch(x, sortx)
		sum += len(sortx) - sx
	}
	return sum
}

func binarySearch(p int, nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if p <= nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
