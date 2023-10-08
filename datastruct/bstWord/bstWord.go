package bstWord

import (
	"strings"
)

const AlphabetSize = 26

type Node struct {
	Children [AlphabetSize]*Node
	IsEnd    bool
}

func New() *Node {
	return &Node{}
}

func (n *Node) Insert(word string) {
	word = sanitize(word)
	current := n

	for _, char := range word {
		index := char - 'a'
		if current.Children[index] == nil {
			current.Children[index] = New()
		}
		current = current.Children[index]
	}

	current.IsEnd = true
}

func (n *Node) Search(word string) bool {
	word = sanitize(word)
	current := n

	for _, char := range word {
		index := char - 'a'
		if current.Children[index] == nil {
			return false
		}
		current = current.Children[index]
	}

	return current.IsEnd
}

func (n *Node) Suggest(prefix string) []string {
	prefix = sanitize(prefix)
	current := n

	for _, char := range prefix {
		index := char - 'a'
		if current.Children[index] == nil {
			return nil
		}
		current = current.Children[index]
	}

	var suggestions []string
	collectWords(current, prefix, &suggestions)
	return suggestions
}

func (n *Node) RemoveWord(word string) {
	n.removeWordRec(word, 0)
}

func (n *Node) removeWordRec(word string, index int) bool {
	if index == len(word) {
		if n.IsEnd {
			n.IsEnd = false
			return len(n.Children) == 0
		}
		return false
	}

	char := word[index]
	if n.Children[char-'a'] == nil {
		return false
	}

	shouldDeleteChild := n.Children[char-'a'].removeWordRec(word, index+1)

	if shouldDeleteChild {
		n.Children[char-'a'] = nil
	}

	return len(n.Children) == 0
}

func collectWords(node *Node, prefix string, suggestions *[]string) {
	if node.IsEnd {
		*suggestions = append(*suggestions, prefix)
	}

	for i, child := range node.Children {
		if child != nil {
			char := rune('a' + i)
			collectWords(child, prefix+string(char), suggestions)
		}
	}
}

func sanitize(word string) string {
	return strings.ToLower(word)
}
