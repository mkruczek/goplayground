package bst_test

import (
	"fmt"
	"some-benchmark/datastruct/bst"
	"testing"
)

func TestNode_Insert(t *testing.T) {

	root := bst.NewNode(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	printTree(root, "***", false)

	if !root.Search(8) {
		t.Errorf("expected to find 8")
	}
}

func TestNode_InsertSameValue_ShouldStoreOnce(t *testing.T) {

	root := bst.NewNode(10)

	root.Insert(10)

	if root.Left != nil {
		t.Errorf("expected left to be nil")
	}
	if root.Right != nil {
		t.Errorf("expected right to be nil")
	}
}

func TestNode_Remove(t *testing.T) {

	root := bst.NewNode(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	root = root.Remove(15)

	if root.Right.Val != 12 {
		t.Errorf("expected right to be 12")
	}
}

func TestNode_InOrder(t *testing.T) {

	root := bst.NewNode(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	got := root.InOrder()

	want := []int{5, 8, 10, 12, 15}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("expected %v, got %v", want, got)
		}
	}
}

func TestNode_Copy(t *testing.T) {
	root := bst.NewNode(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	copied := root.Copy()

	want := root.InOrder()
	got := copied.InOrder()

	for i, v := range got {
		if v != want[i] {
			t.Errorf("expected %v, got %v", want, got)
		}
	}
}

func TestNode_Copy_ShouldNotBeSameInstance(t *testing.T) {
	root := bst.NewNode(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	copied := root.Copy()

	if copied == root {
		t.Errorf("expected to be different instances")
	}

	if copied.Left == root.Left {
		t.Errorf("expected to be different instances")
	}

	if copied.Right == root.Right {
		t.Errorf("expected to be different instances")
	}
}

func TestNode_Clean(t *testing.T) {
	root := bst.NewNode(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	root.Clean()

	if root.Left != nil {
		t.Errorf("expected left to be nil")
	}

	if root.Right != nil {
		t.Errorf("expected right to be nil")
	}

	if root.Val != 0 {
		t.Errorf("expected val to be 0")
	}
}

func printTree(node *bst.Node, prefix string, isLeft bool) {
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

		printTree(node.Left, prefix, true)
		printTree(node.Right, prefix, false)
	}
}
