package main

import (
	"fmt"
	"reflect"
)

func main() {

	assertType(nil)
	assertType(777)
	assertType(true)
	assertType("^_^")

	assertType(make(chan int))
	assertType(make(chan bool))
	assertType(make(chan string))

	assertReflection(nil)
	assertReflection(123)
	assertReflection(false)
	assertReflection("0_0")

	assertReflection(make(chan int))
	assertReflection(make(chan bool))
	assertReflection(make(chan string))

}

// assertType uses a type switch but requires explicitly listing each channel type (chan int, chan bool, etc.).
func assertType(val any) {
	switch val.(type) {
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	case chan int:
		fmt.Println("int channel")
	case chan string:
		fmt.Println("string channel")
	case chan bool:
		fmt.Println("bool channel")
	default:
		fmt.Println("type assertion failed")
	}
}

// assertReflection uses the reflect package to detect types, allowing generic channel recognition.
func assertReflection(val any) {
	if val == nil {
		fmt.Println("type assertion failed")
		return
	}
	t := reflect.TypeOf(val)
	switch t.Kind() {
	case reflect.Int:
		fmt.Println("integer")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("boolean")
	case reflect.Chan:
		fmt.Println("channel")
	default:
		fmt.Println("type assertion failed")
	}
}

/*
Output:
type assertion failed
integer
boolean
string
int channel
bool channel
string channel
type assertion failed
integer
boolean
string
channel
channel
channel
*/
