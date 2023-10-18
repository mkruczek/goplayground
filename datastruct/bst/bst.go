package bst

import (
	"fmt"
	"strings"
)

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

func Max(root *node) int {
	if root == nil {
		return 0
	}

	if root.Right == nil {
		return root.Val
	}

	return Max(root.Right)
}

// AscOrder returns slice of values in ascending order	InOrderTraversal
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

// DestOrder returns slice of values in descending order - !InOrderTraversal
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

// Copy returns a copy of the tree - PreOrderTraversal
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

func (n *node) Mirror() *node {
	if n == nil {
		return nil
	}

	n.Left, n.Right = n.Right, n.Left

	n.Left.Mirror()
	n.Right.Mirror()

	return n
}

// Clean removes all nodes from the tree - PostOrderTraversal
func (n *node) Clean() {
	if n == nil {
		return
	}

	n.Left.Clean()  //left
	n.Right.Clean() //right
	n.Val = 0       //action
	n.Left = nil    //action
	n.Right = nil   //action
}

// LevelOrderTraversal is a BFS implementation
// BFS - Breadth First Search - przeszukiwanie wszerz, wzdluz poziomow
func (n *node) LevelOrderTraversal() [][]int {
	if n == nil {
		return nil
	}

	var result [][]int
	queue := []*node{n}

	for len(queue) > 0 {
		var currentLevel []int
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, current.Val)

			if current.Left != nil {
				queue = append(queue, current.Left)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		result = append(result, currentLevel)
	}

	return result
}

// LowestCommonAncestor znajduje LCA dla dwóch węzłów w BST
func (n *node) LowestCommonAncestor(p, q *node) *node {
	if n == nil {
		return nil
	}

	if n.Val > p.Val && n.Val > q.Val {
		//lca is in left subtree
		return n.Left.LowestCommonAncestor(p, q)
	}

	if n.Val < p.Val && n.Val < q.Val {
		//lca is in right subtree
		return n.Right.LowestCommonAncestor(p, q)
	}

	//lca is n, because n.Val is between p.Val and q.Val
	//  rekurencyjnie jest wywoływany kod który zawsze zwraca jakiegoś noda niezależnie od tego czy jest to lca czy nie
	// az do momentu gdy n.Val jest pomiędzy p.Val i q.Val
	return n
}

// LowestCommonAncestor znajduje LCA dla dwóch węzłów w drzewie binarnym
func (n *node) LowestCommonAncestor2(p, q *node) *node {
	// Kroki bazowe:
	// 1. Jeśli bieżący węzeł jest pusty (nil), zwracamy nil, ponieważ nie ma LCA.
	if n == nil {
		return nil
	}

	// 2. Jeśli wartość bieżącego węzła jest równa wartości któregoś z węzłów p lub q,
	// to bieżący węzeł jest LCA.
	// jako ze node moze byc LCA dla samego siebie
	// czyli w tym kroku rekurencyjnie znalezlismy nasze p lub q
	// które w nastepnym kroku pzypiszemy do leftLCA lub rightLCA
	if n.Val == p.Val || n.Val == q.Val {
		return n
	}

	// 3. Rekurencyjnie znajdujemy  LCA w lewym poddrzewie.
	leftLCA := n.Left.LowestCommonAncestor2(p, q)

	// 4. Rekurencyjnie znajdujemy LCA w prawym poddrzewie.
	rightLCA := n.Right.LowestCommonAncestor2(p, q)

	// 5. Jeśli znajdziemy LCA w obu poddrzewach, to bieżący węzeł jest LCA.
	if leftLCA != nil && rightLCA != nil {
		return n
	}

	// 6. Jeśli LCA jest tylko w lewym poddrzewie, zwracamy go.
	if leftLCA != nil {
		return leftLCA
	}

	// 7. Jeśli LCA jest tylko w prawym poddrzewie, zwracamy go.
	return rightLCA
}

func IsSubTree(root *node, subRoot *node) bool {

	if subRoot == nil {
		return true
	}

	if root == nil {
		return false
	}

	if root.Val == subRoot.Val && IsSameTree(root, subRoot) {
		return true
	}

	return IsSubTree(root.Left, subRoot) || IsSubTree(root.Right, subRoot)
}

func IsSameTree(root *node, subRoot *node) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	if root.Val != subRoot.Val {
		return false
	}

	return IsSameTree(root.Left, subRoot.Left) && IsSameTree(root.Right, subRoot.Right)
}

// means that the trees have the same structure with the possibility of swapping left and right children of a number of nodes
func IsSimilar(a, b *node) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return IsSimilar(a.Left, b.Left) && IsSimilar(a.Right, b.Right) || IsSimilar(a.Left, b.Right) && IsSimilar(a.Right, b.Left)
}

func isBST(root *node) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && root.Left.Val > root.Val {
		return false
	}

	if root.Right != nil && root.Right.Val < root.Val {
		return false
	}

	return isBST(root.Left) && isBST(root.Right)
}

func NodeDistance(root, p, q *node) int {
	lca := root.LowestCommonAncestor2(p, q)
	return distance(lca, p) + distance(lca, q)
}

func distance(root, p *node) int {
	if root == nil {
		return 0
	}

	if root.Val == p.Val {
		return 0
	}

	if root.Val > p.Val {
		return 1 + distance(root.Left, p)
	}

	return 1 + distance(root.Right, p)
}

func FindKSmallestInOrder(root *node, k int) []int {

	if root == nil || k <= 0 {
		return nil
	}
	var result []int

	var inOrder func(root *node) []int
	inOrder = func(root *node) []int {

		if root == nil {
			return nil
		}

		inOrder(root.Left) //goToLeft

		if len(result) < k {
			result = append(result, root.Val) //action
			inOrder(root.Right)               //goToRight
		}

		return result
	}

	inOrder(root)

	//return result[:k]
	return result
}

func FindKBiggstElements(root *node, k int) []int {

	if root == nil || k <= 0 {
		return nil
	}
	var result []int

	var inOrder func(root *node) []int
	inOrder = func(root *node) []int {

		if root == nil {
			return nil
		}

		inOrder(root.Left) //goToLeft

		if len(result) < k {
			result = append(result, root.Val) //action
			inOrder(root.Right)               //goToRight
		}

		return result
	}

	inOrder(root)

	//return result[:k]
	return result
}

func FindKEvenNodsInTheInterval(root *node, k, low, max int) []int {
	if root == nil {
		return nil
	}

	var result []int

	var inOrder func(root *node) []int
	inOrder = func(root *node) []int {
		if root == nil {
			return nil
		}

		inOrder(root.Left) //goToLeft

		if root.Val >= low && root.Val <= max && root.Val%2 == 0 {
			result = append(result, root.Val) //action
		}

		inOrder(root.Right) //goToRight

		return result
	}
	inOrder(root)

	return result[:k]

}

func Serialize(root *node) string {
	if root == nil {
		return ""
	}

	var result []string

	var preOrder func(root *node)
	preOrder = func(root *node) {
		if root == nil {
			result = append(result, "null")
			return
		}

		result = append(result, fmt.Sprintf("%d", root.Val))
		preOrder(root.Left)
		preOrder(root.Right)
	}
	preOrder(root)

	return strings.Join(result, ",")
}

func Deserialize(data string) *node {
	if data == "" {
		return nil
	}

	var nodes []string
	nodes = strings.Split(data, ",")

	var buildTree func() *node
	buildTree = func() *node {
		if len(nodes) == 0 {
			return nil
		}

		val := nodes[0]
		nodes = nodes[1:]

		if val == "null" {
			return nil
		}

		node := &node{}
		fmt.Sscanf(val, "%d", &node.Val)
		node.Left = buildTree()
		node.Right = buildTree()

		return node
	}

	return buildTree()
}

func RebalanceBST(root *node) *node {
	if root == nil {
		return nil
	}

	var nodes []int

	var inOrder func(root *node)
	inOrder = func(root *node) {
		if root == nil {
			return
		}

		inOrder(root.Left)
		nodes = append(nodes, root.Val)
		inOrder(root.Right)
	}
	inOrder(root)

	return buildTree(nodes)
}

func buildTree(nodes []int) *node {
	if len(nodes) == 0 {
		return nil
	}

	mid := len(nodes) / 2

	node := &node{Val: nodes[mid]}
	node.Left = buildTree(nodes[:mid])
	node.Right = buildTree(nodes[mid+1:])

	return node
}

func MaxDepth(root *node) int {
	if root == nil {
		return 0
	}

	return 1 + max(MaxDepth(root.Left), MaxDepth(root.Right))
}

func MinDepth(root *node) int {
	if root == nil {
		return 0
	}

	return 1 + min(MinDepth(root.Left), MinDepth(root.Right))
}

func MaxPathSum(root *node) int {
	if root == nil {
		return 0
	}

	var maxSum int

	var maxPathSum func(root *node) int
	maxPathSum = func(root *node) int {
		if root == nil {
			return 0
		}

		left := max(0, maxPathSum(root.Left))
		right := max(0, maxPathSum(root.Right))

		maxSum = max(maxSum, left+right+root.Val)

		return max(left, right) + root.Val
	}
	maxPathSum(root)

	return maxSum
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
