package dailyExercise

import "math/rand"

type RandomizedSet struct {
	HashMap map[int]int
	Nums    []int
	Idx     int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		HashMap: make(map[int]int),
		Nums:    make([]int, 200000),
		Idx:     0,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.HashMap[val]
	if ok {
		return false
	}
	this.HashMap[val] = this.Idx
	this.Nums[this.Idx] = val
	this.Idx++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.HashMap[val]
	if !ok {
		return false
	}
	delete(this.HashMap, val)
	if index == this.Idx-1 {
		this.HashMap[val] = index
	}
	this.Nums[index] = this.Nums[this.Idx-1]
	this.Idx--
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.Nums[rand.Intn(this.Idx)]
}
