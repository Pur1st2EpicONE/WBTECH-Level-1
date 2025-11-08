package main

import "fmt"

// Human struct with Name and Age fields
type Human struct {
	Name string
	Age  uint8
}

// Action struct embeds Human
// This allows Action to "inherit" all methods of Human
type Action struct {
	Human
}

// GodFunction is a simple constructor for Human
func GodFunction(name string, age uint8) *Human {
	return &Human{Name: name, Age: age}
}

// handleApple is a method on Human
// It prints a message depending on whether Name is set
func (h *Human) handleApple() {
	if h.Name != "" {
		fmt.Printf("%s eats the apple.\n", h.Name)
	} else {
		fmt.Println("No one eats the apple. Life is good.")
	}
}

// handleApple on Action shadows the Human method
// Calling action.handleApple() will invoke this method
func (a Action) handleApple() {
	fmt.Println("The apple is no more. Surely, it's not a big deal.")
}

func main() {

	Adam := GodFunction("Adam", 20)
	Eve := GodFunction("Eve", 20)

	action := new(Action)

	Adam.handleApple()
	Eve.handleApple()

	action.handleApple()
	action.Human.handleApple()

}

/*
Output:
Adam eats the apple.
Eve eats the apple.
The apple is no more. Surely, it's not a big deal.
No one eats the apple. Life is good.
*/
