package bst

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func NewNode(val int) *Node {
	return &Node{Val: val}
}

func (n *Node) Insert(val int) {

	if n.Val == val {
		return
	}

	if val <= n.Val {
		if n.Left == nil {
			n.Left = NewNode(val)
		} else {
			n.Left.Insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = NewNode(val)
		} else {
			n.Right.Insert(val)
		}
	}

}

func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}
	if n.Val == val {
		return true
	}
	if val < n.Val {
		return n.Left.Search(val)
	}
	return n.Right.Search(val)
}

func (n *Node) Remove(val int) *Node {
	if n == nil {
		return nil
	}
	if val < n.Val {
		n.Left = n.Left.Remove(val)
		return n
	}
	if val > n.Val {
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
	smallest := n.Right
	for {
		if smallest != nil && smallest.Left != nil {
			smallest = smallest.Left
		} else {
			break
		}
	}
	n.Val = smallest.Val
	n.Right = n.Right.Remove(n.Val)
	return n
}

func (n *Node) InOrder() []int {
	if n == nil {
		return nil
	}

	var result []int

	result = append(result, n.Left.InOrder()...)
	result = append(result, n.Val)
	result = append(result, n.Right.InOrder()...)

	return result
}

func (n *Node) Copy() *Node {
	if n == nil {
		return nil
	}

	return &Node{
		Val:   n.Val,
		Left:  n.Left.Copy(),
		Right: n.Right.Copy(),
	}
}

func (n *Node) Clean() {
	if n == nil {
		return
	}

	n.Left.Clean()
	n.Right.Clean()
	n.Val = 0
	n.Left = nil
	n.Right = nil
}
