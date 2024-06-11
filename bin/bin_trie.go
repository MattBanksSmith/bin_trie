package bin

import (
	"fmt"
)

type BinData struct {
	isPrepaid bool
}

type TrieNode struct {
	children [10]*TrieNode
	data     BinData
}

type Trie struct {
	root *TrieNode
}

func NewBinRangeTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

func (t *Trie) Insert(bin string, data BinData) {
	node := t.root

	for _, digit := range bin {
		childIndex := digit - '0'
		if node.children[childIndex] == nil {
			node.children[childIndex] = &TrieNode{}
		}
		node = node.children[childIndex]
	}
	node.data = data
}

func (t *Trie) Search(binRange string) BinData {

	node := t.root
	for i, digit := range binRange {
		childIndex := digit - '0'

		//bad data guard
		if childIndex < 0 || childIndex > 10 {
			panic(fmt.Sprintf("bad data %v", childIndex))
		}

		//no child guard
		if node.children[childIndex] == nil {
			return node.data
		}

		//update to next node down
		node = node.children[childIndex]

		//if final character
		if i == len(binRange)-1 {
			return node.data
		}
	}
	return node.data
}
