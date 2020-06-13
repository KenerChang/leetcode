package lrucache

// import (
// 	"fmt"
// )

import (
	"container/list"
)

type LRUCache struct {
	List     *list.List
	Map      map[int]*list.Element
	Capacity int
}

type CacheItem struct {
	Key, Value int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Capacity: capacity,
		Map:      map[int]*list.Element{},
		List:     list.New(),
	}

	return cache
}

func (this *LRUCache) Get(key int) int {
	cacheItem, found := this.Map[key]
	if !found {
		return -1
	}

	this.List.MoveToBack(cacheItem)
	return cacheItem.Value.(*CacheItem).Value
}

func (this *LRUCache) Put(key int, value int) {
	cacheItem, found := this.Map[key]
	if found {
		cacheItem.Value.(*CacheItem).Value = value
		this.List.MoveToBack(cacheItem)
		return
	}

	if this.List.Len() >= this.Capacity {
		target := this.List.Front()
		this.List.Remove(target)

		delete(this.Map, target.Value.(*CacheItem).Key)
	}

	toCache := &CacheItem{
		Key:   key,
		Value: value,
	}

	ele := this.List.PushBack(toCache)
	this.Map[key] = ele
}
