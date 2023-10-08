package bst

import "fmt"

type node struct {
	Val   int
	Left  *node
	Right *node
}

func New(val int) *node {
	return &node{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func (n *node) Insert(val int) {
	if n.Val == val {
		return
	}

	//add to left
	if n.Val > val {
		if n.Left == nil {
			n.Left = &node{Val: val}
			return
		}
		n.Left.Insert(val)
		return
	}

	//add to right
	if n.Right == nil {
		n.Right = &node{Val: val}
		return
	}

	n.Right.Insert(val)
}

func (n *node) Search(val int) bool {

	if n == nil {
		return false
	}

	if n.Val == val {
		return true
	}

	if n.Val > val {
		return n.Left.Search(val)
	}

	return n.Right.Search(val)
}

func (n *node) Remove(val int) *node {

	if n == nil {
		return nil
	}

	if n.Val > val {
		n.Left = n.Left.Remove(val)
		return n
	}

	if n.Val < val {
		n.Right = n.Right.Remove(val)
		return n
	}

	if n.Left == nil && n.Right == nil {
		return nil
	}

	if n.Left == nil {
		return n.Right
	}

	if n.Right == nil {
		return n.Left
	}

	lower := n.Right.Min()
	n.Val = lower.Val
	n.Right = n.Right.Remove(lower.Val)

	return n
}

func (n *node) Min() *node {
	if n == nil {
		return nil
	}

	if n.Left == nil {
		return n
	}

	return n.Left.Min()
}

func (n *node) AscOrder() []int {

	if n == nil {
		return nil
	}

	var result []int

	result = append(result, n.Left.AscOrder()...)
	result = append(result, n.Val)
	result = append(result, n.Right.AscOrder()...)

	return result
}

func (n *node) DestOrder() []int {

	if n == nil {
		return nil
	}

	var result []int

	result = append(result, n.Right.DestOrder()...)
	result = append(result, n.Val)
	result = append(result, n.Left.DestOrder()...)

	return result
}

func (n *node) Copy() *node {
	if n == nil {
		return nil
	}

	return &node{
		Val:   n.Val,
		Left:  n.Left.Copy(),
		Right: n.Right.Copy(),
	}
}

func (n *node) Clean() {
	if n == nil {
		return
	}

	n.Left.Clean()
	n.Right.Clean()
	n.Val = 0
	n.Left = nil
	n.Right = nil
}

func PrintTree(node *node, prefix string, isLeft bool) {
	if node != nil {
		fmt.Printf("%s", prefix)
		if isLeft {
			fmt.Printf("|-- ")
			prefix += "|   "
		} else {
			fmt.Printf("|-- ")
			prefix += "    "
		}

		fmt.Println(node.Val)

		PrintTree(node.Left, prefix, true)
		PrintTree(node.Right, prefix, false)
	}
}
