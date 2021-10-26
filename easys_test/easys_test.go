package main

import (
	"go-daily-problems/easys"
	"testing"
)

func TestAppleTrees(t *testing.T) {
	trees := []int{2, 1, 2, 3, 3, 1, 3, 5}
	max := easys.MaxTreePath(trees)

	if max != 4 {
		t.Fatalf("Test failed %d expected 2", max)
	}
}

func TestToeplitz(t *testing.T) {
	m := [][]int{
		{1, 2, 3, 4, 8},
		{5, 1, 2, 3, 4},
		{4, 5, 1, 2, 3},
		{7, 4, 5, 1, 2},
	}

	hasToeplitz := easys.IsToeplitz(m)
	if hasToeplitz == false {
		t.Fatal("Toeplitz should be true")
	}
}
