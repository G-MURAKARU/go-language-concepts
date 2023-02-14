package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// defer - invoke a function to run at a later point in execution, LIFO/stack
	// panic - program reaches a state where it can no longer continue running
	// recover - save a panicked programme

	// notice LIFO behaviour of defer
	// deferred functions take the argument values at the point where the function is deferred
	// not at the end of the function
	a := "first"
	defer fmt.Println(a)
	a = "not first"
	fmt.Println("second")
	defer fmt.Println("third")

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
		// Fatal() = Print() + os.Exit()
	}
	// example of defer use case - not to forget to close resources
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
