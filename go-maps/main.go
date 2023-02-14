package main

import (
	"fmt"
)

func main() {
	// maps are like dictionaries or hash tables - they are UNORDERED
	// keys have to be 'equality-enabled' e.g. hashable data types cannot be used as keys
	// you can't use a map as a key
	// declaring a map: map[key-data-type]value-data-type
	grades := map[string]int{
		"math": 73,
		"eng":  95,
		"swa":  90,
		"sci":  100,
		"sst":  80,
	}
	fmt.Println(grades)

	// other:
	otherGrades := make(map[string]int)
	fmt.Println(otherGrades)

	// adding values
	grades["cre"] = 95
	fmt.Println(grades)

	// retrieving values (note: will work even if key does not exist - defaults to zero-value of value-data-type)
	fmt.Println("math grade:", grades["math"])

	// to check if the key exists, use comma-ok -> ok returns false if it does not exist
	if val, ok := grades["chem"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("unavailable")
	}

	// deleting values
	delete(grades, "swa")
	fmt.Println(grades)

	// maps are also REFERENCE DATA TYPES
	j := grades
	fmt.Printf("%p, %p (same)\n", j, grades)
}
