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

func TestGetRelativePeak(t *testing.T){
   peak, err := easys.GetRelativePeak([]int{1,4,5,2,10,2})
   if peak != 5 && peak != 10 && err != nil {
        t.Fatalf("%d is not a peak\n", peak)
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
