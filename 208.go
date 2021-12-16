package main

import "fmt"

type Trie struct {
	root treeNode
}

type treeNode struct {
	element byte
	toPtr   map[byte]*treeNode
	isEnd   bool
}

func Constructor() Trie {
	return Trie{root: treeNode{' ', make(map[byte]*treeNode), false}}
}

func (this *Trie) Insert(word string) {
	now := &this.root //now 这里是上级节点
	if len(word) == 1 {
		b := byte(word[0])
		if now.toPtr[byte(b)] == nil {
			now.toPtr[byte(b)] = &treeNode{byte(b), make(map[byte]*treeNode), true}
		} else {
			now.toPtr[byte(b)].isEnd = true
		}
	}
	for k, b := range word {
		if now.toPtr[byte(b)] == nil {
			now.toPtr[byte(b)] = &treeNode{byte(b), make(map[byte]*treeNode), k == len(word)-1}
		}
		now = now.toPtr[byte(b)] //
		if k == len(word)-1 {
			now.isEnd = true
		}
	}
}

func (this *Trie) Search(word string) bool {
	now := this.root.toPtr[byte(word[0])]
	if now == nil {
		return false
	}
	for k, v := range word[0:] {
		if k == len(word)-1 {
			return now.element == byte(v) && now.isEnd
		}
		if now.element != byte(v) || now.toPtr[byte(word[k+1])] == nil {
			return false
		}
		now = now.toPtr[byte(word[k+1])]
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	now := this.root.toPtr[byte(prefix[0])]
	if now == nil {
		return false
	}
	for k, v := range prefix[0:] {
		if k == len(prefix)-1 {
			return now.element == byte(v)
		}
		if now.element != byte(v) || now.toPtr[byte(prefix[k+1])] == nil {
			return false
		}
		now = now.toPtr[byte(prefix[k+1])]
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("b")
	fmt.Println(trie.Search("apple")) // return True
	fmt.Println(trie.Search("bdsad"))
	fmt.Println(trie.Search("appjklfsdjflk")) // return False
	fmt.Println(trie.StartsWith("app"))       // return True
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // return True

}
