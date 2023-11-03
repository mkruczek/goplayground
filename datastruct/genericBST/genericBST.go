package genericBST

type Hashable interface {
	Hash() int
}

type Node[T Hashable] struct {
	Val   T
	Left  *Node[T]
	Right *Node[T]
}

func NewNode[T Hashable](val T) *Node[T] {
	return &Node[T]{Val: val}
}

func (n *Node[T]) Insert(val T) {
	if n == nil {
		return
	}

	//we have this value already
	if n.Val.Hash() == val.Hash() {
		return
	}

	//add to left
	if n.Val.Hash() > val.Hash() {
		if n.Left == nil {
			n.Left = &Node[T]{Val: val}
			return
		}
		n.Left.Insert(val)
		return
	}
}

func (n *Node[T]) IsSameTree(t *Node[T]) bool {

	if n == nil && t == nil {
		return true
	}

	if n == nil || t == nil {
		return false
	}

	if n.Val.Hash() != t.Val.Hash() {
		return false
	}

	return n.Left.IsSameTree(t.Left) && n.Right.IsSameTree(t.Right)
}
