package main

import (
	"fmt"
)

func main() {
	// declaring an array - arrays can only store variables of one data type
	// arrays have a FIXED SIZE that MUST BE DEFINED at COMPILE TIME
	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v, (%d grades)\n", grades, len(grades))

	// other:
	var myArray [3]uint8
	fmt.Println("my array:", myArray)

	// creating a matrix
	var identityMatrix = [3][3]int{
		// no need to precede each with [3]int, redundant as it has already been defined above
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Printf("identity matrix: %v, %T\n", identityMatrix, identityMatrix)

	// arrays in Go are VALUE TYPES, unlike REFERENCE TYPES in other languages
	otherIdentityMatrix := identityMatrix // creates a COPY of the whole array,
	// instead of making it point to the same underlying array

	fmt.Printf("first: %p, second: %p (different memory locations)\n", &identityMatrix, &otherIdentityMatrix)

	// use pointers to avoid copying the whole array
	identityPointer := &identityMatrix
	fmt.Println(*identityPointer)
}
