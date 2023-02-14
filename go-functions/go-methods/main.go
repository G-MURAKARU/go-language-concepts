// METHOD SETS - value receivers vs pointer receivers
// you can assign a method to any data type, though mostly structs

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type shape interface {
	area() float64
	perimeter() float64
}

func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2.0)
}

func (sq square) area() float64 {
	return math.Pow(sq.length, 2.0)
}

func (c circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}

func (sq square) perimeter() float64 {
	return sq.length * 4
}

func info(s shape) {
	fmt.Println("area:", s.area(), "perimeter:", s.perimeter())
}

func main() {
	my_circle := circle{
		radius: 14.5,
	}
	my_square := square{
		length: 6.9,
	}

	my_other_square := square{
		length: 4.20,
	}
	/*
		NOTE:
		since circle and square are both receivers of the area() method,
		(peep area() methods are identical in declaration: name, parameters and return(s))
		then they both inplement the shape interface, i.e. they are both
		"of type shape"
	*/

	/*
		METHOD SETS:
		the method set of a defined type T consists of all methods
		defined with a receiver of type T
		the method set of a pointer to a type T consists of all methods
		defined with a receiver of type T or *T i.e.
		var1 T -> func (arg T) funcname() {...}
		var2 *T -> func (arg T | arg *T) funcname() {...}
		OR
		func (arg T) funcname() {...} -> var1 T | var2 *T
		func (arg *T) funcname() {...} -> var2 *T
	*/

	/*
		the circle area method, defined with a receiver of type *T,
		can only be called on an identifier of type *T
	*/
	info(&my_circle) // identifier of type *T, receiver of type *T (area) and type T (perimeter)
	/*
		NOTE, BUT THIS RUNS FINE:
		my_circle.area()
	*/

	/*
		the square area method, defined with a receiver of type T,
		can be called on identifiers of both type T and *T
	*/
	info(my_square)        // identifier of type T, receivers of type T (both area and perimeter)
	info(&my_other_square) // identifier of type *T, receivers of type T (both area and perimeter)
}
