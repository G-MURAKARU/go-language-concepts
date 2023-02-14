package main

import (
	"fmt"
	"log"
)

// PANIC - RECOVER DEMO
// remember - call stack
// NB: panics are handled higher up the call stack

func panicker() {
	fmt.Println("about to panic")

	// recover() is only called in deferred functions
	defer func() {
		if err := recover(); err != nil {
			// in here you're saying YOU'RE going to deal with it
			log.Println("Error:", err)
			// if there is no way to deal with reasons for panic, REPANIC
			// <panic(some error message)>
		}
	}()
	panic("something bad happened")
	// notice: code below is unreachable
	// <fmt.Println("done panicking")>
}

func main() {
	fmt.Println("start")
	panicker()
	// notice: the application returns to a normal state of execution (the statement below executes)
	// if no call to recover() panic would reach go runtime and program will exit (below statement won't execute)
	fmt.Println("end")
}
