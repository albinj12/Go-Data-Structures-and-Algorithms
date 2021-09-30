package main

import (
	"fmt"
)

type node struct {
	children map[string]*node
	wordEnd  bool
}

type trie struct {
	root *node
}

func main() {

	t := trie{root: &node{
		children: make(map[string]*node),
	}}

}

func (t *trie) insert(word string) {
	ptr := t.root
	for _, v := range word {
		if ptr.children[string(v)] == nil {
			ptr.children[string(v)] = &node{
				children: map[string]*node{},
			}
		}

		ptr = ptr.children[string(v)]
	}
	ptr.wordEnd = true

}

func (t *trie) search(word string) bool {
	ptr := t.root
	for _, v := range word {
		if ptr.children[string(v)] == nil {
			return false
		}
		ptr = ptr.children[string(v)]
	}
	if ptr.wordEnd == true {
		return true
	}

	return false
}

func (t *trie) delete(word string) bool {
	deleteNode(t.root, word, 0)
	return true
}

func deleteNode(trieNode *node, word string, index int) *node {

	if index == len(word) {

		if trieNode.wordEnd {
			trieNode.wordEnd = false
		}

		if len(trieNode.children) == 0 {
			trieNode = nil
		}
		return trieNode
	}
	char := string(word[index])
	if trieNode.children[char] != nil {
		index++
		trieNode.children[char] = deleteNode(trieNode.children[char], word, index)
	}

	if trieNode.children[char] == nil {
		delete(trieNode.children, char)

	}

	if len(trieNode.children) == 0 && trieNode.wordEnd == false {
		trieNode = nil
	}

	return trieNode
}

func (t *trie) displayWords(trieNode *node, word string) {
	if trieNode.wordEnd == true {
		fmt.Println(word)
	}
	for key, value := range trieNode.children {
		t.displayWords(value, word+string(key))
	}
}
