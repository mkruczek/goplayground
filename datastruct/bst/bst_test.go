package bst

import (
	"reflect"
	"slices"
	"testing"
)

func TestNode_Insert(t *testing.T) {

	root := New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	PrintTree(root, "***", false)

	//should find
	if !root.Search(8) {
		t.Errorf("expected to find 8")
	}

	//should not find
	if root.Search(81) {
		t.Errorf("expected to not find 81")
	}
}

func TestNode_InsertSameValue_ShouldStoreOnce(t *testing.T) {

	root := New(10)

	root.Insert(10)

	if root.left != nil {
		t.Errorf("expected left to be nil")
	}
	if root.right != nil {
		t.Errorf("expected right to be nil")
	}
}

func TestNode_Remove(t *testing.T) {

	root := New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(13)
	root.Insert(8)
	root.Insert(12)

	root.Remove(5)

	PrintTree(root, "***", false)

	if root.right.left.val != 13 {
		t.Errorf("expected right to be 12")
	}
}

func TestNode_Remove_Root(t *testing.T) {

	root := New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(13)
	root.Insert(8)
	root.Insert(12)

	root.Remove(10)

	PrintTree(root, "***", false)

	if root.right.left.val != 13 {
		t.Errorf("expected right to be 12")
	}
}

func TestNode_Min(t *testing.T) {

	root := New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)
	root.Insert(1)

	mi := root.Min().val

	if mi != 1 {
		t.Errorf("expected min to be 5")
	}
}

func TestNode_Max(t *testing.T) {

	root := New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)
	root.Insert(1)

	mi := root.Max().val

	if mi != 15 {
		t.Errorf("expected min to be 5")
	}
}

func TestNode_InOrder(t *testing.T) {

	root := New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	got := root.AscOrder()

	want := []int{5, 8, 10, 12, 15}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("expected %v, got %v", want, got)
		}
	}
}

func TestNode_DestOrder(t *testing.T) {

	root := New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	got := root.DestOrder()

	want := []int{15, 12, 10, 8, 5}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("expected %v, got %v", want, got)
		}
	}
}

func TestNode_Copy(t *testing.T) {
	root := New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	copied := root.Copy()

	want := root.AscOrder()
	got := copied.AscOrder()

	if !slices.Equal(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}

}

func TestNode_Mirror(t *testing.T) {
	root := New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	root.Mirror()

	want := []int{15, 12, 10, 8, 5}
	got := root.AscOrder()

	for i, v := range got {
		if v != want[i] {
			t.Errorf("expected %v, got %v", want, got)
		}
	}
}

func TestNode_Copy_ShouldNotBeSameInstance(t *testing.T) {
	root := New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	copied := root.Copy()

	if copied == root {
		t.Errorf("expected to be different instances")
	}

	if copied.left == root.left {
		t.Errorf("expected to be different instances")
	}

	if copied.right == root.right {
		t.Errorf("expected to be different instances")
	}
}

func TestNode_Clean(t *testing.T) {
	root := New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	root.Clean()

	if root.left != nil {
		t.Errorf("expected left to be nil")
	}

	if root.right != nil {
		t.Errorf("expected right to be nil")
	}

	if root.val != 0 {
		t.Errorf("expected val to be 0")
	}
}

func TestNode_LevelOrderTraversal(t *testing.T) {
	root := New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)
	root.Insert(3)
	root.Insert(1)

	want := [][]int{{10}, {5, 15}, {3, 8, 12}, {1}}
	got := root.LevelOrderTraversal()

	for i, v := range got {
		if !slices.Equal(v, want[i]) {
			t.Errorf("expected %v, got %v", want, got)
		}
	}
}

func TestNode_LowestCommonAncestor(t *testing.T) {
	root := New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)
	root.Insert(3)
	root.Insert(1)

	lca := root.LowestCommonAncestor2(New(1), New(12))

	if lca.val != 10 {
		t.Errorf("expected lca to be 10")
	}
}

func TestMaxValOnPath(t *testing.T) {
	// test tree
	//        10
	//       /  \
	//      5    15
	//     / \   / \
	//    3   7 12  18
	root := &node{
		val: 10,
		left: &node{
			val:   5,
			left:  &node{val: 3},
			right: &node{val: 7},
		},
		right: &node{
			val:   15,
			left:  &node{val: 12},
			right: &node{val: 18},
		},
	}

	expected := 18
	result := root.Max().val
	if result != expected {
		t.Errorf("Oczekiwano %d, otrzymano %d", expected, result)
	}

	singleNode := &node{val: 42}
	result = singleNode.Max().val
	expected = 42
	if result != expected {
		t.Errorf("Oczekiwano %d, otrzymano %d", expected, result)
	}
}

func TestMinDepth(t *testing.T) {
	// testTree
	//        10
	//       /  \
	//      5    15
	//     / \   / \
	//    3   7 12  18
	//             /
	//			  17
	root := &node{
		val: 10,
		left: &node{
			val:   5,
			left:  &node{val: 3},
			right: &node{val: 7},
		},
		right: &node{
			val:   15,
			left:  &node{val: 12},
			right: &node{val: 18, left: &node{val: 17}},
		},
	}

	expected := 3
	result := MinDepth(root)
	if result != expected {
		t.Errorf("Oczekiwano %d, otrzymano %d", expected, result)
	}

	expected = 4
	result = MaxDepth(root)
	if result != expected {
		t.Errorf("Oczekiwano %d, otrzymano %d", expected, result)
	}

	result = MinDepth(nil)
	expected = 0
	if result != expected {
		t.Errorf("Oczekiwano %d, otrzymano %d", expected, result)
	}

	singleNode := &node{val: 42}
	result = MinDepth(singleNode)
	expected = 1
	if result != expected {
		t.Errorf("Oczekiwano %d, otrzymano %d", expected, result)
	}
}

func TestFindKSmallestInOrder(t *testing.T) {
	tests := []struct {
		data     []int
		k        int
		expected []int
	}{
		{nil, 5, nil},
		{[]int{4, 2, 6, 1, 3, 5, 7}, 3, []int{1, 2, 3}},
		{[]int{10}, 1, []int{10}},
		{[]int{8, 5, 12, 3, 7, 10, 14, 1, 4, 6, 9, 11, 13, 15}, 14, []int{1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
	}

	for i, test := range tests {
		root := createBST(test.data)
		result := FindKSmallestInOrder(root, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d failed. Expected: %v, Got: %v", i+1, test.expected, result)
		}
	}
}

func createBST(data []int) *node {
	var root *node
	for _, val := range data {
		root = insertBST(root, val)
	}
	return root
}

func insertBST(root *node, val int) *node {
	if root == nil {
		return &node{val: val}
	}
	if val < root.val {
		root.left = insertBST(root.left, val)
	} else if val > root.val {
		root.right = insertBST(root.right, val)
	}
	return root
}

func TestIsSubTree(t *testing.T) {
	tests := []struct {
		data     []int
		subtree  []int
		expected bool
	}{
		{nil, nil, true},
		{[]int{4, 2, 6, 1, 3, 5, 7}, []int{2, 1, 3}, true},
		{[]int{4, 2, 6, 1, 3, 5, 7}, []int{2, 1, 3, 4}, false},
		{[]int{4, 2, 6, 1, 3, 5, 7}, []int{2, 1, 3, 5}, false},
		{[]int{4, 2, 6, 1, 3, 5, 7}, []int{2, 1, 3, 5, 6}, false},
		{[]int{4, 2, 6, 1, 3, 5, 7}, []int{2, 1, 3, 5, 6, 7}, true},
		{[]int{4, 2, 6, 1, 3, 5, 7}, []int{2, 1, 3, 5, 6, 7, 8}, false},
		{[]int{4, 2, 6, 1, 3, 5, 7}, []int{2, 1, 3, 5, 6, 7, 8, 9}, false},
		{[]int{4, 2, 6, 1, 3, 5, 7}, []int{2, 1, 3, 5, 6, 7, 8, 9, 10}, false},
		{[]int{3, 4, 5, 1, 2}, []int{4, 1, 2}, true},
	}

	for i, test := range tests {
		root := createBST(test.data)
		subtree := createBST(test.subtree)
		result := IsSubTree(root, subtree)
		if result != test.expected {
			t.Errorf("Test %d failed. Expected: %v, Got: %v", i+1, test.expected, result)
		}
	}
}

func TestIsSimilar(t *testing.T) {
	tests := []struct {
		data1    []int
		data2    []int
		expected bool
	}{
		{nil, nil, true},
		{[]int{4, 2, 6, 1, 3, 5, 7}, []int{4, 2, 6, 1, 3, 5, 7}, true},
		{[]int{4, 2, 6, 1, 3, 5, 7}, []int{4, 2, 6, 1, 3, 5}, false},
		{[]int{4, 2, 6, 1, 3, 5}, []int{4, 2, 6, 1, 3, 5, 7}, false},
		{[]int{4, 6, 2, 1, 3, 6, 7}, []int{4, 2, 6, 1, 3, 6, 7}, true},
		{[]int{4, 2, 6, 1, 3, 5, 7}, []int{4, 2, 6, 1, 4, 5, 7}, false},
	}

	for i, test := range tests {
		root1 := createBST(test.data1)
		root2 := createBST(test.data2)
		result := IsSimilar(root1, root2)
		if result != test.expected {
			t.Errorf("Test %d failed. Expected: %v, Got %v", i+1, test.expected, result)
		}
	}
}
