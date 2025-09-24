package main

import "fmt"

func main() {

	A := []int{1, 2, 2, 3, 4, 3, 1, 5, 1, 2, 1, 1, 3}
	B := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(makeSet(A))
	fmt.Println(makeSet(B))

}

// makeSet returns a slice containing the unique elements of the input slice.
// T is a generic type constrained by comparable, so it can be used as a map key.
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

/*
Output:
[1 2 3 4 5]
[cat dog tree]
*/
