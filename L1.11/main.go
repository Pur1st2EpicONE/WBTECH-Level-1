package main

import "fmt"

func main() {

	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	C := []string{"Желание", "Ржавый", "Семнадцать", "Рассвет", "Печь", "Один", "Товарный вагон"}
	D := []string{"Рассвет", "Семнадцать", "Печь", "Возвращение на Родину", "Ржавый", "Один"}

	fmt.Println(intersection(A, B))
	fmt.Println(intersection(C, D))

}

// intersection finds all unique elements present in both input slices.
// T is a type parameter constrained by "comparable", so it can be used as a map key.
func intersection[T comparable](A, B []T) []T {
	hm := make(map[T]struct{})
	for _, val := range A {
		hm[val] = struct{}{}
	}
	res := []T{}
	for _, val := range B {
		if _, found := hm[val]; found {
			res = append(res, val)
			delete(hm, val)
		}
	}
	return res
}

/*
Output:
[2 3]
[Рассвет Семнадцать Печь Ржавый Один]
*/
