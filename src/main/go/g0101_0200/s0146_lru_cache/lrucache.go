package s0146_lru_cache

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Hash_Table #Design #Linked_List
// #Doubly_Linked_List #Udemy_Linked_List #Big_O_Time_O(1)_Space_O(capacity)
// #2024_03_21_Time_421_ms_(93.60%)_Space_69_MB_(91.02%)

import "container/list"

type KeyValue struct {
	key int
	val int
}

type LRUCache struct {
	m        map[int]*list.Element
	q        *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{make(map[int]*list.Element), list.New(), capacity}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.m[key]; ok {
		this.q.MoveToFront(n)
		return n.Value.(KeyValue).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	kv := KeyValue{key, value}

	if n, ok := this.m[key]; ok {
		n.Value = kv
		this.q.MoveToFront(n)
		this.m[key] = n
	} else {
		this.m[key] = this.q.PushFront(kv)
	}

	if len(this.m) > this.capacity {
		delete(this.m, this.q.Back().Value.(KeyValue).key)
		this.q.Remove(this.q.Back())
	}
}

/*
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
