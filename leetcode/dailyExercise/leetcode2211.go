package dailyExercise

import (
	"strings"
)

func countCollisions(directions string) int {
	res := 0
	s := strings.ReplaceAll(directions, "RL", "S")
	res += 2 * (len(directions) - len(s))
	s = strings.TrimLeft(s, "L")
	s = strings.TrimRight(s, "R")
	return res + len(s) - strings.Count(s, "S")
}
