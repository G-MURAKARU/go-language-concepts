package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main() {
	// decalring an int variable
	var a int

	// declaring a pointer to an integer
	var b *int

	// assignment
	a = 42
	// assigning b as a pointer to the memory location where a's value is stored
	b = &a

	// voila
	fmt.Println("a (value):", a, "\nb (value - a memory location):", b, "\nb(dereferenced - showing what is stored in b):", *b)

	var ms *myStruct

	ms = &myStruct{foo: 128}

	// dot operator has a higher precedence than the dereferencing operator, therefore enclosed * in parens e.g. (*ms).foo
	fmt.Println(ms.foo)

	ms2 := new(myStruct)

	// this is sick but okay.
	ms2.foo = 42

	fmt.Println(ms2.foo)
}
