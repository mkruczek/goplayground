package bstWord

import (
	"slices"
	"testing"
)

func Test_main(t *testing.T) {
	words := []string{"apple", "banana", "cherry", "date", "fig"}
	root := &Node{}

	for _, word := range words {
		root.Insert(word)
	}

	if got, want := root.Search("apple"), true; got != want {
		t.Errorf("expected %v, got %v", want, got)
	}

	if got, want := root.Search("banana"), true; got != want {
		t.Errorf("expected %v, got %v", want, got)
	}

	root.RemoveWord("banana")
	if got, want := root.Search("banana"), false; got != want {
		t.Errorf("expected %v, got %v", want, got)
	}

	root.Insert("abc")
	root.Insert("abcd")
	root.Insert("abcde")

	if got, want := root.Search("abc"), true; got != want {
		t.Errorf("expected %v, got %v", want, got)
	}

	want := []string{"abc", "abcd", "abcde"}
	got := root.Suggest("abc")

	if !slices.Equal(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}

}
