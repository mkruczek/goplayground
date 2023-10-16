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

//func (n *Node) Suggest(prefix string) []string {
//	prefix = sanitize(prefix)
//	current := n
//
//	for _, char := range prefix {
//		index := char - 'a'
//		if current.Children[index] == nil {
//			return nil
//		}
//		current = current.Children[index]
//	}
//
//	var suggestions []string
//	collectWords(current, prefix, &suggestions)
//	return suggestions
//}
//
//func collectWords(node *Node, prefix string, suggestions *[]string) {
//	if node.IsEnd {
//		*suggestions = append(*suggestions, prefix)
//	}
//
//	for i, child := range node.Children {
//		if child != nil {
//			char := rune('a' + i)
//			collectWords(child, prefix+string(char), suggestions)
//		}
//	}
//}

func (n *Node) Suggest(prefix string) []string {
	prefix = sanitize(prefix)
	current := n

	//szuakmy noda na ktorym konczy sie prefix
	for _, char := range prefix {
		index := char - 'a'
		if current.Children[index] == nil {
			return nil
		}
		current = current.Children[index]
	}

	var result []string
	var suggest func(node *Node, currentPrefix string)
	suggest = func(node *Node, currentPrefix string) {
		if node == nil {
			return
		}

		//dla prefixu "abc" , jesli abc jest slowem to dodajemy do result "abc"
		if node.IsEnd {
			result = append(result, prefix+currentPrefix)
		}

		//dla prefixu "abc" , jesli abc jest czescia slowa np abcdefg to wiedzac ze jestemy na nodzie "c" to sprawdzamy wszystkie dzieci wezla "c"
		// wiemy ze "w baziie" jest abcdefg, wiec gdy dojdziemy do dziecka "d" to wywolujemy suggest(child, currentPrefix(pusty bo abc jest mainprefixem)+"d") i sprawdzamy dzieci wezla "d"
		for i, child := range node.Children {
			char := rune('a' + i)
			suggest(child, currentPrefix+string(char))
		}
	}

	//w tym miejscu current wskazuje na ostatni node z prefixu czyli dla "abc" na node "c"
	//currentPrefix jest pusty bo nie ma wiecej liter po "abc", wiec wywołujemy suggest(current, "") i sprawdzamy dzieci wezla "c"
	suggest(current, "")

	return result
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

func sanitize(word string) string {
	return strings.ToLower(word)
}
