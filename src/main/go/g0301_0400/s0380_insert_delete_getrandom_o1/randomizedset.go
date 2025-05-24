package s0380_insert_delete_getrandom_o1

// #Medium #Array #Hash_Table #Math #Design #Randomized #Programming_Skills_II_Day_20
// #Top_Interview_150_Array/String #2025_05_24_Time_47_ms_(87.65%)_Space_47.12_MB_(83.31%)

import (
	"math/rand"
)

type RandomizedSet struct {
	list []int
	dict map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		list: make([]int, 0),
		dict: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.dict[val]; exists {
		return false
	}
	this.dict[val] = len(this.list)
	this.list = append(this.list, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, exists := this.dict[val]; !exists {
		return false
	}
	last := len(this.list) - 1
	lastVal := this.list[last]
	idx := this.dict[val]
	this.list[idx] = lastVal
	this.dict[lastVal] = idx
	this.list = this.list[:last]
	delete(this.dict, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))] // NOSONAR
}
