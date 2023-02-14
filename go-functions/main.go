package main

import "fmt"

/*
	GO FUNCTION SIGNATURE:
	func (receiver) funcName(paremeters) (return type) {...}
	- receiver is optional -> when a func has a receiver, it becomes a method (think OOP python)
	- you can have value receivers and pointer receivers
	- parameters are optional, included if needed
	- parens around return type are only necessary if there are multiple returns or named returns
*/

func sum(values ...int) *int {
	var result int
	for _, v := range values {
		result += v
	}

	// in Go, you can return a pointer to a variable that was initialised locally
	// the Go compiler will promote the variable from the function's local stack to the heap

	//advantage - the variable can be worked with as a true value (bypassing dereferencing) within the function
	return &result
}

func mul(values ...int) (result int) {
	// Go allows us to have named returns
	// notice: do not have to declare the return variable
	result = 1
	for _, v := range values {
		result *= v
	}
	// as seen here, no need to specify what is being returned
	return
}

func div(a float32, b float32) (float32, error) {
	// multiple returns in Go (see above)
	// IDIOMATIC GO: returning an error value in a function that could return an error
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	s := sum(1, 2, 3, 4, 5)
	m := mul(1, 2, 3, 4, 5)
	fmt.Println("The sum is:", *s)
	fmt.Println("The product is:", m)

	// some idiomatic go :)
	d, err := div(1.8, 9.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The quotient is: %.2f", d)
	}

	// anonymous functions / closures
	func() {
		fmt.Println("i'm an anoymous func")
	}() // parens invoke the function - has to be invoked

	// functions as variables
	myFunc := func() {
		fmt.Println("my func.")
	}
	myFunc()

	// wild demo of anonymous functions
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		}
		return a / b, nil
	}

	dd, err := divide(1.8, 4.5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dd)
	}
}
