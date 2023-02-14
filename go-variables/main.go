package main

import (
	"fmt"
	"strconv"
)

// declaring package-level variables, can't use `walrus` operator
var myGlobalFloat float32

// declaring multiple variables simultaneously
var (
	// can be considered in design to group related variables
	name         string  = "Gicheru"
	age          int     = 24
	heightMeters float32 = 1.77
	isMale       bool    = true
)

// dummy var to demo redeclaration and shadowing
var dummy int = 24

func main() {
	// declaring a variable using the `var` keyword
	// used to declare a variable but not initialise it yet
	var myInt int
	// REMEMBER: declared but uninitialised variables default to the ZERO-VALUE of its DATA TYPE
	myInt = 24

	// on one line
	// used if value does not give enough info for Go compiler to infer its type,
	// e.g. declaring what looks like an int as a float32
	var myOtherInt float32 = 25

	// declaring a variable using the `walrus` operator
	yetAnotherInt := 26

	fmt.Println(myInt)
	fmt.Println(myOtherInt)
	fmt.Println(yetAnotherInt)

	// demonstration of redeclaration
	var dummy int = 30
	fmt.Printf("%v, %T -> Locally defined variable takes precedence - SHADOWING (global var is shadowed by local var)\n", dummy, dummy)

	// type conversions
	// Go cannot do implicit type conversions
	converted := float32(dummy)
	fmt.Printf("%f, %T\n", converted, converted)

	// note: strings in Go work similarly to strings in C
	// therefore, converting an int to a string will treat the int as the ascii value of the string
	i := 42 // string(i) returns "*" instead of "42"

	// fix: import and use strconv.Itoa()
	j := strconv.Itoa(i)
	fmt.Printf("CORRECT: ascii-unicode value of %d = *, Sprint converted value of %d (%T) = %s (%T)\n", i, i, i, j, j)
}
