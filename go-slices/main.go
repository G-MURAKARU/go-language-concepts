package main

import (
	"fmt"
)

func main() {
	// declaring a slice (note: nothing in the [])
	a := []byte{1, 2, 3, 4, 5}
	fmt.Printf("%v, %T\n", a, a)

	// slices are based on arrays, therefore the length of the slice differs from the capacity of the slice i.e, length of the underlying array
	// first arg defines slice type
	// second arg defines slice length
	// third (optional) arg defines the slice capacity/underlying array length, defaults to value in second arg
	b := make([]int, 10, 20) // initialised with all 0's
	fmt.Println("old slice length:", len(b), "old array length:", cap(b))
	// but you can't index beyond the slice i.e. b[10] will error

	// if input exceeds slice capacity, go compiler initialises another array with double the capacity of the previous array
	// go garbage collector gets rid of the old array
	for i := 1; i < 16; i++ {
		b = append(b, i)
	}
	fmt.Println("new slice length:", len(b), "new array length:", cap(b))

	// slices are REFERENCE TYPES, unlike arrays (a slice points to the underlying array)
	c := a
	fmt.Printf("first: %p, second: %p (they are the same)\n", a, c)

	// SLICE OPERATIONS - use slicing and appending where applicable
	// removing first element
	d := a[1:]
	fmt.Println(d)
	// removing last element
	e := a[:len(a)-1]
	fmt.Println(e)
	// deleting an element at an index e.g index 2
	f := append(a[:2], a[3:]...)
	fmt.Println("f:", f, "a:", a)
}
