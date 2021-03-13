package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

type P struct {
	L int
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

// Type P implement method M.
// P implicit implement of interface I.
// P could be declared in interface I.
func (p P) M() {
	fmt.Println("-----> ", p.L)
}

func main() {
	var i I = T{"hello"}
	i.M()

	var r I = P{5}
	r.M()
}
