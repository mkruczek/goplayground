package bst

import (
	"fmt"
	"strings"
)

type node struct {
	val   int
	left  *node
	right *node
}

func New(val int) *node {
	return &node{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func (n *node) Insert(val int) {
	if n.val == val {
		return
	}

	//add to left
	if n.val > val {
		if n.left == nil {
			n.left = &node{val: val}
			return
		}
		n.left.Insert(val)
		return
	}

	//add to right
	if n.right == nil {
		n.right = &node{val: val}
		return
	}

	n.right.Insert(val)
}

func (n *node) Search(val int) bool {

	if n == nil {
		return false
	}

	if n.val == val {
		return true
	}

	if n.val > val {
		return n.left.Search(val)
	}

	return n.right.Search(val)
}

func (n *node) Remove(val int) *node {

	if n == nil {
		return nil
	}

	if n.val > val {
		n.left = n.left.Remove(val)
		return n
	}

	if n.val < val {
		n.right = n.right.Remove(val)
		return n
	}

	if n.left == nil && n.right == nil {
		return nil
	}

	if n.left == nil {
		return n.right
	}

	if n.right == nil {
		return n.left
	}

	lower := n.right.Min()
	n.val = lower.val
	n.right = n.right.Remove(lower.val)

	return n
}

func (n *node) Min() *node {
	if n == nil {
		return nil
	}

	if n.left == nil {
		return n
	}

	return n.left.Min()
}

func (n *node) Max() *node {
	if n == nil {
		return nil
	}

	if n.right == nil {
		return n
	}

	return n.right.Max()
}

// AscOrder returns slice of values in ascending order	InOrderTraversal
func (n *node) AscOrder() []int {

	if n == nil {
		return nil
	}

	var result []int

	result = append(result, n.left.AscOrder()...)
	result = append(result, n.val)
	result = append(result, n.right.AscOrder()...)

	return result
}

// DestOrder returns slice of values in descending order - !InOrderTraversal
func (n *node) DestOrder() []int {

	if n == nil {
		return nil
	}

	var result []int

	result = append(result, n.right.DestOrder()...)
	result = append(result, n.val)
	result = append(result, n.left.DestOrder()...)

	return result
}

// Copy returns a copy of the tree - PreOrderTraversal
func (n *node) Copy() *node {
	if n == nil {
		return nil
	}

	return &node{
		val:   n.val,
		left:  n.left.Copy(),
		right: n.right.Copy(),
	}
}

func (n *node) Mirror() *node {
	if n == nil {
		return nil
	}

	n.left, n.right = n.right, n.left

	n.left.Mirror()
	n.right.Mirror()

	return n
}

// Clean removes all nodes from the tree - PostOrderTraversal
func (n *node) Clean() {
	if n == nil {
		return
	}

	n.left.Clean()  //left
	n.right.Clean() //right
	n.val = 0       //action
	n.left = nil    //action
	n.right = nil   //action
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
			currentLevel = append(currentLevel, current.val)

			if current.left != nil {
				queue = append(queue, current.left)
			}

			if current.right != nil {
				queue = append(queue, current.right)
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

	if n.val > p.val && n.val > q.val {
		//lca is in left subtree
		return n.left.LowestCommonAncestor(p, q)
	}

	if n.val < p.val && n.val < q.val {
		//lca is in right subtree
		return n.right.LowestCommonAncestor(p, q)
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
	if n.val == p.val || n.val == q.val {
		return n
	}

	// 3. Rekurencyjnie znajdujemy  LCA w lewym poddrzewie.
	leftLCA := n.left.LowestCommonAncestor2(p, q)

	// 4. Rekurencyjnie znajdujemy LCA w prawym poddrzewie.
	rightLCA := n.right.LowestCommonAncestor2(p, q)

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

	if root.val == subRoot.val && IsSameTree(root, subRoot) {
		return true
	}

	return IsSubTree(root.left, subRoot) || IsSubTree(root.right, subRoot)
}

func IsSameTree(root *node, subRoot *node) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	if root.val != subRoot.val {
		return false
	}

	return IsSameTree(root.left, subRoot.left) && IsSameTree(root.right, subRoot.right)
}

// means that the trees have the same structure with the possibility of swapping left and right children of a number of nodes
func IsSimilar(a, b *node) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return IsSimilar(a.left, b.left) && IsSimilar(a.right, b.right) || IsSimilar(a.left, b.right) && IsSimilar(a.right, b.left)
}

func isBST(root *node) bool {
	if root == nil {
		return true
	}

	if root.left != nil && root.left.val > root.val {
		return false
	}

	if root.right != nil && root.right.val < root.val {
		return false
	}

	return isBST(root.left) && isBST(root.right)
}

func NodeDistance(root, p, q *node) int {
	lca := root.LowestCommonAncestor2(p, q)
	return distance(lca, p) + distance(lca, q)
}

func distance(root, p *node) int {
	if root == nil {
		return 0
	}

	if root.val == p.val {
		return 0
	}

	if root.val > p.val {
		return 1 + distance(root.left, p)
	}

	return 1 + distance(root.right, p)
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

		inOrder(root.left) //goToLeft

		if len(result) < k {
			result = append(result, root.val) //action
			inOrder(root.right)               //goToRight
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

		inOrder(root.left) //goToLeft

		if len(result) < k {
			result = append(result, root.val) //action
			inOrder(root.right)               //goToRight
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

		inOrder(root.left) //goToLeft

		if root.val >= low && root.val <= max && root.val%2 == 0 {
			result = append(result, root.val) //action
		}

		inOrder(root.right) //goToRight

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

		result = append(result, fmt.Sprintf("%d", root.val))
		preOrder(root.left)
		preOrder(root.right)
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
		fmt.Sscanf(val, "%d", &node.val)
		node.left = buildTree()
		node.right = buildTree()

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

		inOrder(root.left)
		nodes = append(nodes, root.val)
		inOrder(root.right)
	}
	inOrder(root)

	return buildTree(nodes)
}

func buildTree(nodes []int) *node {
	if len(nodes) == 0 {
		return nil
	}

	mid := len(nodes) / 2

	node := &node{val: nodes[mid]}
	node.left = buildTree(nodes[:mid])
	node.right = buildTree(nodes[mid+1:])

	return node
}

func MaxDepth(root *node) int {
	if root == nil {
		return 0
	}

	return 1 + max(MaxDepth(root.left), MaxDepth(root.right))
}

func MinDepth(root *node) int {
	if root == nil {
		return 0
	}

	return 1 + min(MinDepth(root.left), MinDepth(root.right))
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

		left := max(0, maxPathSum(root.left))
		right := max(0, maxPathSum(root.right))

		maxSum = max(maxSum, left+right+root.val)

		return max(left, right) + root.val
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

		fmt.Println(node.val)

		PrintTree(node.left, prefix, true)
		PrintTree(node.right, prefix, false)
	}
}
