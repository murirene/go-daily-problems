package easys

import "math"

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
