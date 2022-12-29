package collections

/**
 * @author 2416144794@qq.com
 * @date 2022/12/29 11:35
 */

type UnionFindByMap struct {
	nums           []int
	numToParentMap map[int]int
}

func NewUnionFindByMap(nums []int) *UnionFindByMap {
	numToParentMap := make(map[int]int)
	for _, num := range nums {
		numToParentMap[num] = num
	}
	return &UnionFindByMap{
		nums:           nums,
		numToParentMap: numToParentMap,
	}
}

func (u *UnionFindByMap) Find(x int) int {
	if x == u.numToParentMap[x] {
		return x
	}
	return u.Find(u.numToParentMap[x])
}

func (u *UnionFindByMap) Union(x, y int) {
	px, py := u.Find(x), u.Find(y)
	if px == py {
		return
	}
	u.numToParentMap[px] = py
}

func (u *UnionFindByMap) GetMaxHerdSize() int {
	resMap := make(map[int]int)
	res := 0
	for num, _ := range u.numToParentMap {
		pnum := u.Find(num)
		resMap[pnum]++
		if res < resMap[pnum] {
			res = resMap[pnum]
		}
	}
	return res
}

func (u *UnionFindByMap) Count() int {
	res := 0
	for num, _ := range u.numToParentMap {
		if num == u.Find(num) {
			res++
		}
	}
	return res
}
