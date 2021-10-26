package easys

import (
	"fmt"
	"math"
)

func MaxTreePath(trees []int) int {
	max := 0
	start := 0
	end := 0
	treeMap := make(map[int]int, 0)

	for end < len(trees) {
		if len(treeMap) <= 2 {
			if _, ok := treeMap[trees[end]]; ok == false {
				treeMap[trees[end]] = 1
			} else {
				treeMap[trees[end]] += 1
			}

			// shrink the path to include 2 trees
			for len(treeMap) > 2 {
				val := trees[start]
				treeMap[val] -= 1

				if treeMap[val] == 0 {
					delete(treeMap, val)
				}
				start++
			}
		}

		max = int(math.Max(float64(max), float64(end-start+1)))
		end++
	}

	return max
}

func resetMatrix(m [][]int) {
	for i, _ := range m {
		for j, _ := range m[i] {
			m[i][j] = -1
		}
	}
}

func IsToeplitz(m [][]int) bool {
	cache := make([][]int, len(m))
	for i := range cache {
		cache[i] = make([]int, len(m[i]))
	}
	resetMatrix(cache)
	d1 := helperToeplitzD1(m, 0, 0, cache)
	resetMatrix(cache)
	d2 := helperToeplitzD2(m, 0, 0, cache)
	fmt.Printf("d1=%v and d2=%v\n", d1, d2)
	return d1 || d2
}

// Implementation with no caching but with both diag checked
func helperToeplitz(m [][]int, i, j int) (bool, bool) {
	d1 := true
	d2 := true

	if i == -1 || j == -1 || i == len(m) || j == len(m[0]) {
		return d1, d2
	}

	// out of bounds returns true
	d1Adj, d2Adj := helperToeplitz(m, i, j+1)
	d1Bot, d2Bot := helperToeplitz(m, i+1, j)

	if i < (len(m)-1) && j < (len(m[0])-1) {
		d1 = m[i][j] == m[i+1][j+1]
	}

	d1 = d1 && d1Adj && d1Bot

	if i > 0 && j < (len(m[0])-1) {
		d2 = m[i][j] == m[i-1][j+1]
	}

	d2 = d2 && d2Adj && d2Bot

	return d1, d2
}

// Top Left to bottom right Toeplitz
func helperToeplitzD1(m [][]int, i, j int, cache [][]int) bool {
	d1 := true
	d1Bot := true
	d1Adj := true

	if i == -1 || j == -1 || i == len(m) || j == len(m[0]) {
		return d1
	}

	if j < (len(cache[0])-1) && cache[1][j+1] == -1 {
		d1Adj = helperToeplitzD1(m, i, j+1, cache)
	} else if j < (len(cache[0]) - 1) {
		fmt.Printf("%d %d cache hit\n", i, j)
		d1Adj = cache[i][j+1] == 1
	}

	if i < (len(cache)-1) && cache[i+1][j] == -1 {
		d1Bot = helperToeplitzD1(m, i+1, j, cache)
	} else if i < (len(cache) - 1) {
		fmt.Printf("%d %d cache hit\n", i, j)
		d1Bot = cache[i+1][j] == 1
	}

	if i < (len(m)-1) && j < (len(m[0])-1) {
		d1 = m[i][j] == m[i+1][j+1]
	}

	return d1 && d1Adj && d1Bot
}

// Top Right to Bottom Left Toeplitz
func helperToeplitzD2(m [][]int, i, j int, cache [][]int) bool {
	d2 := true
	d2Adj := true
	d2Bot := true

	if i == -1 || j == -1 || i == len(m) || j == len(m[0]) {
		return d2
	}

	if j < (len(cache[0])-1) && cache[i][j+1] == -1 {
		d2Adj = helperToeplitzD2(m, i, j+1, cache)
	} else if j < (len(cache[0]) - 1) {
		fmt.Printf("%d %d cache hit\n", i, j)
		d2Adj = cache[i][j+1] == 1
	}

	if i < (len(m)-1) && cache[i+1][j] == -1 {
		d2Bot = helperToeplitzD2(m, i+1, j, cache)
	} else if i < (len(m) - 1) {
		fmt.Printf("%d %d cache hit\n", i, j)
		d2Bot = cache[i+1][j] == 1
	}

	if i > 0 && j < (len(m[0])-1) {
		d2 = m[i][j] == m[i-1][j+1]
	}

	d2 = d2 && d2Adj && d2Bot
	if d2 == true {
		cache[i][j] = 1
	} else {
		cache[i][j] = 0
	}

	return d2
}
