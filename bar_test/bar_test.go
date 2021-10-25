package main 

import (
    "testing"
    "go-daily-problems/bar"
)

func TestHelloWorld(t *testing.T) {
    preferences := make(map[int][]int, 0)
    preferences[0] = []int{0,1,3,6}
    preferences[1] = []int{1,4,7}
    preferences[2] = []int{2,4,7,5}
    preferences[3] = []int{1,3,2,5}
    preferences[4] = []int{5,8}
    drinks := bar.GetMinRecipes(preferences)

    if drinks != 2 {
	    t.Fatalf("shouldbe 2 not %d", drinks)
     }
}
