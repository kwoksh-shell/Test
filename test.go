package main

import (
	"container/list"
)

type node struct {
	key, value, freq int
}

type LFUCache struct {
	capacity, size, minFreq int
	keyMap                  map[int]*list.Element
	freqMap                 map[int]*list.List
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		size:     0,
		minFreq:  0,
		keyMap:   map[int]*list.Element{},
		freqMap:  map[int]*list.List{},
	}
}

func (this *LFUCache) Get(key int) int {
	elem, exist := this.keyMap[key]
	if !exist {
		return -1
	}
	n := elem.Value.(*node)
	if this.freqMap[n.freq].Len() == 0 && this.minFreq == n.freq {
		this.minFreq++
	}
	n.freq++
	if this.freqMap[n.freq] == nil {
		this.freqMap[n.freq] = list.New()
	}
	newElem := this.freqMap[n.freq].PushBack(n)
	this.keyMap[key] = newElem
	return n.value
}

func (this *LFUCache) Put(key, value int) {
	if this.capacity <= 0 {
		return
	}
	if elem, exist := this.keyMap[key]; exist {
		n := elem.Value.(*node)
		this.freqMap[n.freq].Remove(elem)
		if this.freqMap[n.freq].Len() == 0 && this.minFreq == n.freq {
			this.minFreq++
		}
		n.freq++
		if this.freqMap[n.freq] == nil {
			this.freqMap[n.freq] = list.New()
		}
		newElem := this.freqMap[n.freq].PushBack(n)
		this.keyMap[key] = newElem
		return
	}
	if this.size >= this.capacity {
		victimLists := this.freqMap[this.minFreq]
		victimElem := victimLists.Front()
		victimNode := victimElem.Value.(*node)
		delete(this.keyMap, victimNode.key)
		victimLists.Remove(victimElem)
		if victimLists.Len() == 0 {
			delete(this.freqMap, this.minFreq)
		}
	}
	newNode := &node{key: key, value: value, freq: 1}
	if this.freqMap[1] == nil {
		this.freqMap[1] = list.New()
	}
	Elem := this.freqMap[1].PushBack(newNode)
	this.keyMap[key] = Elem
	this.minFreq = 1
	this.size++
}
