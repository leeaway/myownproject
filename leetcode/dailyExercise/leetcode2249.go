package dailyExercise

import "fmt"

func countLatticePoints(circles [][]int) int {
	resMap := make(map[string]bool)
	for _, circle := range circles {
		x, y, r := circle[0], circle[1], circle[2]
		for i := x - r; i <= x+r; i++ {
			for j := y - r; j <= y+r; j++ {
				key := arrayToString(i, j)
				if !resMap[key] && isInCircle(circle, []int{x, y}) {
					resMap[key] = true
				}
			}
		}
	}
	return len(resMap)
}

func isInCircle(circle, point []int) bool {
	return circle[2]*circle[2] >= (circle[0]-point[0])*(circle[0]-point[0])+(circle[1]-point[1])*(circle[1]-point[1])
}

func arrayToString(x, y int) string {
	return fmt.Sprintf("%v:%v", x, y)
}
