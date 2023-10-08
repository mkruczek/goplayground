package bst_test

import (
	"some-benchmark/datastruct/bst"
	"testing"
)

func TestNode_Insert(t *testing.T) {

	root := bst.New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	bst.PrintTree(root, "***", false)

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

	root := bst.New(10)

	root.Insert(10)

	if root.Left != nil {
		t.Errorf("expected left to be nil")
	}
	if root.Right != nil {
		t.Errorf("expected right to be nil")
	}
}

func TestNode_Remove(t *testing.T) {

	root := bst.New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(13)
	root.Insert(8)
	root.Insert(12)

	root.Remove(5)

	bst.PrintTree(root, "***", false)

	if root.Right.Left.Val != 13 {
		t.Errorf("expected right to be 12")
	}
}

func TestNode_Remove_Root(t *testing.T) {

	root := bst.New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(13)
	root.Insert(8)
	root.Insert(12)

	root.Remove(10)

	bst.PrintTree(root, "***", false)

	if root.Right.Left.Val != 13 {
		t.Errorf("expected right to be 12")
	}
}

func TestNode_Min(t *testing.T) {

	root := bst.New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)
	root.Insert(1)

	mi := root.Min().Val

	if mi != 1 {
		t.Errorf("expected min to be 5")
	}
}

func TestNode_InOrder(t *testing.T) {

	root := bst.New(10)

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

	root := bst.New(10)

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
	root := bst.New(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)

	copied := root.Copy()

	want := root.AscOrder()
	got := copied.AscOrder()

	for i, v := range got {
		if v != want[i] {
			t.Errorf("expected %v, got %v", want, got)
		}
	}
}

func TestNode_Copy_ShouldNotBeSameInstance(t *testing.T) {
	root := bst.New(10)

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
	root := bst.New(10)

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
