package main

import "fmt"

func main() {

	A := []int{1, 2, 2, 3, 4, 3, 1, 5, 1, 2, 1, 1, 3}
	B := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(makeSet(A))
	fmt.Println(makeSet2(B))

}

// makeSet returns a slice of unique elements from the input slice.
// Fast but unordered — uses map for deduplication with two passes.
func makeSet[T comparable](arr []T) []T {
	hm := make(map[T]struct{})
	for _, val := range arr {
		hm[val] = struct{}{}
	}
	set := make([]T, 0, len(hm))
	for key := range hm {
		set = append(set, key)
	}
	return set
}

// makeSet2 returns a slice of unique elements from the input slice.
// Preserves order of first occurrence — single pass with conditional checks.
func makeSet2[T comparable](arr []T) []T {
	hm := make(map[T]struct{})
	res := []T{}
	for _, val := range arr {
		if _, ok := hm[val]; !ok {
			hm[val] = struct{}{}
			res = append(res, val)
		}
	}
	return res
}

/*
Output:
[1 2 3 4 5]
[cat dog tree]
*/
