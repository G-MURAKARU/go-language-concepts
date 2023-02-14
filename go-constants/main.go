package main

import (
	"fmt"
)

const a int8 = 10

// enumerated constants with iota
const (
	// iota used for incrementing by 1, can be evaluated
	z = iota + 5
	// the compiler can infer the pattern, no need to call iota each time
	y
	x
	w
)

const (
	// compiler sees this as different const block, resets iota
	_ = iota // returns 0 so unused (right-only value)
	// these can be evaluated at compile time, so they work
	kb = 1 << (iota * 10)
	mb
	gb
	tb
)

// using iota + bit shifting to set flags in a single byte
const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	// declaring constants - can be set to any primitive type
	const myConst int = 42
	fmt.Println(myConst)

	// constants have to be set to a value that is calculable at compile time
	// i.e. it can't be set to sth that is determined at runtime e.g. the return of a function
	// e.g. <const otherConst float64 = math.Sin(90)> will fail

	// constants cannot be set to inherently mutable types e.g. arrays, slices, maps, etc
	// e.g. <const my_ints []int := {1, 2, 3}> will fail

	// constants can be redeclared and shadowed too
	const a int = 20
	fmt.Printf("shadowed const: %d, %T\n", a, a)

	// constants can be evaluated with variables, if they are of the same type
	var b int = 30
	fmt.Println("a + b (const + var) =", a+b)

	// constants can be untyped - type is evaluated by compiler
	// compiler optimisation behind constants - compiler searches for where const name is used and LITERALLY replaces it with that value
	// think const and volatile data types from C
	// therefore, const type is evaluated at runtime
	const c = 6
	var d uint8 = 5
	fmt.Printf("c + d (untyped const [%T] + typed var [%T]) = %v [%T]\n", c, d, c+d, c+d)

	var e float32 = 3.14159
	fmt.Printf("c + e (untyped const [%T] + typed var [%T]) = %v [%T]\n", c, e, c+e, c+e)

	fmt.Println("\niota (enumerated constants) demo:")
	fmt.Printf("%v, %v, %v, %v\n\n", z, y, x, w)
	fmt.Println("iota reset:")
	fmt.Printf("kb: %v\t\t\t%b\n", kb, kb)
	fmt.Printf("mb: %v\t\t\t%b\n", mb, mb)
	fmt.Printf("gb: %v\t\t\t%b\n", gb, gb)
	fmt.Printf("tb: %v\t\t%b\n", tb, tb)

	// using flags to set roles of a certain user
	// OR bitwise operator is used to 'combine' permissions, sets bits
	fmt.Println("using iota + bit shifting to set certain flags (kind of like registers in embedded):")
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope // = 00100101
	fmt.Printf("%b\n", roles)

	// to check if the person has certain permissions - BITMASKING
	if allowed := roles&canSeeFinancials == canSeeFinancials; allowed {
		fmt.Println("You can see the financials.")
	}

	fmt.Printf("At headquarters? %v\n", isHeadquarters&roles == isHeadquarters)
}
