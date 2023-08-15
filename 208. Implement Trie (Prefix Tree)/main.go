package main

type Node struct {
	val      byte
	isWord   bool
	elements [26]*Node
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{elements: [26]*Node{}},
	}
}

func (t *Trie) Insert(word string) {
	temp := t.root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'

		if temp.elements[idx] == nil {
			temp.elements[idx] = &Node{
				val:      word[i],
				elements: [26]*Node{},
			}
		}

		if i == len(word)-1 {
			temp.elements[idx].isWord = true
		}

		temp = temp.elements[idx]
	}
}

func (t *Trie) Search(word string) bool {
	temp := t.root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'

		if temp.elements[idx] == nil {
			return false
		}

		temp = temp.elements[idx]
	}

	return temp.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	temp := t.root
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'

		if temp.elements[idx] == nil {
			return false
		}

		temp = temp.elements[idx]
	}

	return true
}
