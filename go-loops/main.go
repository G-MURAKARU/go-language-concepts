package main

import "fmt"

func main() {
	// LOOPS
	// note: go does not have a while loop
	// it does have break and continue, like python

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Printf("%d %d\n", i, j)
	}
	k := 10 // function scope

	// notice the first semicolon
	fmt.Println("first k")
	for ; k < 16; k++ {
		fmt.Println(k)
	}

	// how go achieves while loops
	// notice no semicolons
	fmt.Println("second k")
	for k < 21 {
		fmt.Println(k)
		k++
	}

	// infinite loop
	l := 0
	for {
		if l == 10 {
			break
		}
		l++
	}

	// you can use a LABEL to break out of a NESTED FOR LOOP
	// place the label where you want the break statement to lead to
MyLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
			if i > 2 && j > 3 {
				break MyLoop
			}
		}
	}

	// to loop over a collection with indeterminate size, use for range loop
	// works like enumerate() in python
	someSlice := []int{11, 44, 64, 54, 23}
	for idx, val := range someSlice {
		fmt.Println("index:", idx, "value:", val)
	}
	// looping over maps is similar to looping over dictionaries in python
}
