package main

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

type LRUCache struct {
	Cap   int
	Table map[int]*Node
	Head  *Node
	Tail  *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{Val: -1}, &Node{Val: -1}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
		Cap:   capacity,
		Table: make(map[int]*Node),
		Head:  head,
		Tail:  tail,
	}
}

func (lru *LRUCache) Get(key int) int {
	if node, has := lru.Table[key]; has {
		lru.unLink(key)
		lru.moveToHead(node)
		return node.Val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if node, has := lru.Table[key]; has {
		node.Val = value
		lru.unLink(key)
		lru.moveToHead(node)
	} else {
		newNode := &Node{Key: key, Val: value}
		if len(lru.Table) == lru.Cap {
			delKey := lru.removeTail()
			delete(lru.Table, delKey)
		}
		lru.Table[key] = newNode
		lru.moveToHead(newNode)
	}
}

func (lru *LRUCache) moveToHead(node *Node) {
	head := lru.Head.Next
	lru.Head.Next = node
	node.Next = head
	node.Pre = lru.Head
	head.Pre = node
}

func (lru *LRUCache) removeTail() int {
	tail := lru.Tail.Pre
	pre := tail.Pre
	pre.Next = lru.Tail
	lru.Tail.Pre = pre
	tail.Pre = nil
	tail.Next = nil
	return tail.Key
}

func (lru *LRUCache) unLink(i int) {
	node := lru.Table[i]
	pre := node.Pre
	next := node.Next
	pre.Next = next
	next.Pre = pre
}
