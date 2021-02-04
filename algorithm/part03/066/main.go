package main

import (
	"container/list"
)

func main() {

}

type Node struct {
	key int
	val int
}

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	link     *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		link:     list.New(),
	}
}

func (s *LRUCache) Get(key int) int {
	v, ok := s.cache[key]
	if ok {
		s.link.MoveToFront(v)
		return v.Value.(*Node).val
	}
	return -1
}

func (s *LRUCache) Put(key int, value int) {
	v, ok := s.cache[key]
	if ok {
		v.Value.(*Node).val = value
		s.link.MoveToFront(v)
	} else {
		if len(s.cache)+1 > s.capacity {
			ele := s.link.Back()
			delete(s.cache, ele.Value.(*Node).key)
			s.link.Remove(ele)
		}
		ele := s.link.PushFront(&Node{
			key: key,
			val: value,
		})
		s.cache[key] = ele
	}
}
