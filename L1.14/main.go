package main

import (
	"fmt"
	"reflect"
)

func main() {

	testTypes()

}

// testTypes calls assertType and assertReflection on various values including basic types,
// bidirectional channels, send-only channels, and receive-only channels.
func testTypes() {

	assertType(nil)
	assertType(777)
	assertType(true)
	assertType("^_^")

	assertType(make(chan int))
	assertType(make(chan bool))
	assertType(make(chan string))

	assertType(make(chan<- int))
	assertType(make(chan<- bool))
	assertType(make(chan<- string))

	assertType(make(<-chan int))
	assertType(make(<-chan bool))
	assertType(make(<-chan string))

	assertReflection(nil)
	assertReflection(123)
	assertReflection(false)
	assertReflection("0_0")

	assertReflection(make(chan int))
	assertReflection(make(chan bool))
	assertReflection(make(chan string))

	assertReflection(make(chan<- int))
	assertReflection(make(chan<- bool))
	assertReflection(make(chan<- string))

	assertReflection(make(<-chan int))
	assertReflection(make(<-chan bool))
	assertReflection(make(<-chan string))

}

// assertType prints the type of val using a type switch.
// It distinguishes between basic types and channels of different directions (bidirectional, send-only, receive-only).
func assertType(val any) {

	switch val.(type) {
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	case chan int:
		fmt.Println("int channel (bidirectional)")
	case chan bool:
		fmt.Println("bool channel (bidirectional)")
	case chan string:
		fmt.Println("string channel (bidirectional)")
	case chan<- int:
		fmt.Println("int channel (send-only)")
	case chan<- bool:
		fmt.Println("bool channel (send-only)")
	case chan<- string:
		fmt.Println("string channel (send-only)")
	case <-chan int:
		fmt.Println("int channel (receive-only)")
	case <-chan bool:
		fmt.Println("bool channel (receive-only)")
	case <-chan string:
		fmt.Println("string channel (receive-only)")
	default:
		fmt.Println("type assertion failed")
	}

}

// assertReflection prints the type of val using reflection.
// It can detect basic types and any channel type, but does not distinguish channel direction.
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
int channel (bidirectional)
bool channel (bidirectional)
string channel (bidirectional)
int channel (send-only)
bool channel (send-only)
string channel (send-only)
int channel (receive-only)
bool channel (receive-only)
string channel (receive-only)
type assertion failed
integer
boolean
string
channel
channel
channel
channel
channel
channel
channel
channel
channel
*/
