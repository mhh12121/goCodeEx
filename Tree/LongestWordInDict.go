package main

import "strings"

/*
Given a list of strings words representing an English Dictionary, find the longest word in words that can be built one character at a time by other words in words. If there is more than one possible answer, return the longest word with the smallest lexicographical order.

If there is no answer, return the empty string.
Example 1:
Input:
words = ["w","wo","wor","worl", "world"]
Output: "world"
Explanation:
The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".
Example 2:
Input:
words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
Output: "apple"
Explanation:
Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".
*/
type NodeWord struct {
	c        byte
	children map[byte]*NodeWord
	end      int
}
type TrieII struct {
	root  *NodeWord
	words []string
}

func (trie *TrieII) insert(word string, index int) {
	cur := trie.root
	for _, value := range []byte(word) { //#todo
		_, ok := cur.children[value]
		if !ok {
			cur.children[value] = &NodeWord{c: value, children: make(map[byte]*NodeWord)}

		}
		cur = cur.children[value]
	}
	cur.end = index
}
func (trie *TrieII) dfs() string {
	ans := ""
	stack := make([]*NodeWord, 0)
	stack = append(stack, trie.root)
	for len(stack) != 0 {
		tmpNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if tmpNode.end > 0 || tmpNode == trie.root {
			if tmpNode != trie.root {
				word := trie.words[tmpNode.end-1]
				if len(word) > len(ans) || len(word) == len(ans) && strings.Compare(word, ans) < 0 {
					ans = word
				}
			}
			for _, value := range tmpNode.children {
				stack = append(stack, value)
			}
		}

	}
	return ans
}
func longestWord(words []string) string {
	trie := &TrieII{root: &NodeWord{c: '0', children: make(map[byte]*NodeWord)}}
	index := 0
	for _, value := range words {
		index++
		trie.insert(value, index)
	}
	trie.words = words
	return trie.dfs()

}
