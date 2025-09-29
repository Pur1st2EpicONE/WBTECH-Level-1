package main

import "fmt"

// oldDevice represents devices that can connect via micro USB
type oldDevice interface {
	connectMicroUSB()
}

// oldPhone is an implementation of oldDevice
type oldPhone struct{}

// connectMicroUSB simulates connecting an old phone with micro USB
func (o *oldPhone) connectMicroUSB() {
	fmt.Println("Connected through micro USB")
}

// newDevice represents devices that can connect via type-C
type newDevice interface {
	connectTypeC()
}

// newPhone is an implementation of newDevice
type newPhone struct{}

// connectTypeC simulates connecting a new phone with type-C
func (n *newPhone) connectTypeC() {
	fmt.Println("Connected through type-C")
}

// adapter makes an oldDevice compatible with the newDevice interface
type adapter struct {
	oldDevice oldDevice
}

// connectTypeC allows an oldDevice to be used where a newDevice is expected
func (a *adapter) connectTypeC() {
	fmt.Println("Using adapter to connect old phone through micro USB")
	a.oldDevice.connectMicroUSB()
}

func main() {

	oldPhone := &oldPhone{}
	oldPhone.connectMicroUSB()

	newPhone := &newPhone{}
	newPhone.connectTypeC()

	var adapter newDevice = &adapter{oldDevice: oldPhone}
	adapter.connectTypeC()

}

/*
Output:
Connected through micro USB
Connected through type-C
Using adapter to connect old phone through micro USB
Connected through micro USB
*/
