package dailyExercise

func fallingSquares(positions [][]int) []int {
	n := len(positions)
	squareHeight := make([][]int, n)
	var res []int
	maxHeight := 0
	for i, pos := range positions {
		squareHeight[i] = make([]int, 3)
		squareHeight[i][0] = pos[0]
		squareHeight[i][1] = pos[1]
		squareHeight[i][2] = pos[1]
		curHeight := calCurHeight(i, squareHeight)
		squareHeight[i][2] = curHeight
		if maxHeight < curHeight {
			maxHeight = curHeight
		}
		res = append(res, maxHeight)
	}
	return res
}

func calCurHeight(index int, squareHeight [][]int) int {
	startHeight := squareHeight[index][2]
	maxHeight := startHeight
	for i := 0; i < index; i++ {
		if isCovered(squareHeight[i], squareHeight[index]) {
			curHeight := startHeight + squareHeight[i][2]
			if curHeight > maxHeight {
				maxHeight = curHeight
			}
		}
	}
	return maxHeight
}

func isCovered(pos1, pos2 []int) bool {
	xs1, xe1 := pos1[0], pos1[0]+pos1[1]
	xs2, xe2 := pos2[0], pos2[0]+pos2[1]
	return !(xe2 <= xs1 || xe1 <= xs2)
}
