package collections

type UnionFind struct {
	Parent []int
	Num    int
}

/*
	初始化时：每个节点的Parent指向自己
*/
func NewUnionFind(n int) *UnionFind {
	people := make([]int, n)
	for i := range people {
		people[i] = i
	}
	return &UnionFind{
		Parent: people,
		Num:    n,
	}
}

func (u *UnionFind) Find(x int) int {
	//父节点为自身，就是找到了
	if u.Parent[x] == x {
		return x
	}
	return u.Find(u.Parent[x])
}

func (u *UnionFind) Union(x, y int) {
	px, py := u.Find(x), u.Find(y)
	if px == py {
		return
	}
	u.Parent[px] = py
}
