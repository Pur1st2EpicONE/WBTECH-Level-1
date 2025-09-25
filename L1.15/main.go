package main

var justString string

func main() {

	someFunc()
	fixedFunc()

}

// someFunc creates a huge string and assigns a substring to justString.
// The substring shares the same underlying array as the original string,
// so it keeps a reference to the full 1KB string in memory.
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

// createHugeString allocates a string of the given size filled with zeros.
func createHugeString(size int) string {
	return string(make([]byte, size))
}

// fixedFunc creates a huge string and copies only the first 100 bytes
// into justString. This avoids keeping a reference to the full 1KB string,
// which saves memory.
func fixedFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}
