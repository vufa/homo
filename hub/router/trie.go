package router

import "sync"

// Trie topic tree of *common.SinkSubs
type Trie struct {
	root *node
	sync.RWMutex
}

// NewTrie creates a trie
func NewTrie() *Trie {
	return &Trie{
		root: newNode(),
	}
}
