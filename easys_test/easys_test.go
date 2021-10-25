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
