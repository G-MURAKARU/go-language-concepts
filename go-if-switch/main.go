package main

import "fmt"

func main() {
	guess := -5

	// IF STATEMENTS
	// short-circuiting ifs
	// if one condition returns true for chained logical OR operations, others will be ignored (axiom of boolean algebra)
	// notice that returnTrue() will not run ("returning true" will not be printed)
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("'returning true' was not printed - short circuited OR")
	}

	// same for logical AND, if any condition returns false, others will be ignored
	// notice that returnTrue() will not run ("returning true" will not be printed)
	if guess >= 1 && returnTrue() && guess <= 100 {
		fmt.Print("lol")
	} else {
		fmt.Println("'returning true' was not printed - short circuited AND")
	}

	// SWITCH STATEMENTS
	// can switch over a range
	// tagged switch - cases must be unique and non-overlapping
	switch i := 2 + 9; i {
	case 1, 2, 3, 4:
		fmt.Println("one - four")
	case 5, 6, 7, 8:
		fmt.Println(" five -  eight")
	default:
		fmt.Println("others")
	}

	// tagless - cases can overlap, first one will run
	j := 10
	switch {
	case j <= 10:
		fmt.Println("<= 10")
		// in order to fall through to the next case, use fallthrough
		// note: fallthrough is logic-less i.e. next case will run automatically
		fallthrough
	case j == 20:
		fmt.Println("<= 20")
	default:
		fmt.Println("> 20")
	}

	// type switching - you can switch on the TYPE of the tag value
	var i interface{} = map[string]uint8{
		"one": 1,
		"two": 2,
	}
	i = 1
	switch i.(type) {
	case int:
		fmt.Println("int")
		// to leave the case early, use break
		if true {
			break
		}
		fmt.Println("still here")
	case float32:
		fmt.Println("float32")
	case string:
		fmt.Println("string")
	case [3]int:
		fmt.Println("[3]int")
	case map[string]uint8:
		fmt.Println("map[string]uint8")
	default:
		fmt.Println("some other type")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
