package dailyExercise

type RecentCounter struct {
	Que []int
}

func Constructor2() RecentCounter {
	return RecentCounter{Que: []int{}}
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.Que) > 0 && t-this.Que[0] > 3000 {
		this.Que = this.Que[1:]
	}
	this.Que = append(this.Que, t)
	return len(this.Que)
}
