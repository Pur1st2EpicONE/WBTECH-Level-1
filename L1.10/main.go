package main

import "fmt"

func main() {

	temps := [8]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(group(temps))

}

// group assigns temperatures to 10-degree ranges by dividing each value by 10,
// converting to int (truncating toward zero), and multiplying back by 10 to get the range start.
// Each range start is used as a key in the map, with a slice of temperatures as the value.
func group(temps [8]float64) map[int][]float64 {
	hm := make(map[int][]float64, 4)
	for _, val := range temps {
		temp := int(val/10) * 10
		hm[temp] = append(hm[temp], val)
	}
	return hm
}

/*
Output:
map[-20:[-25.4 -27 -21] 10:[13 19 15.5] 20:[24.5] 30:[32.5]]
*/
