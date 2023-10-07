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
