package main

import (
	"net/http"
)

// panics should only be used when an application is in a state where it CANNOT or SHOULD NOT continue to execute
// e.g. divide by zero

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	// example of error, the port is blocked/already in use
	// ListenAndServe just returns an error value (it is reasonable that you tried to access a blocked port)
	// therefore, it just returns an error - it tells us that things did not go as expected
	// NO RESPONSE IS A VALID RESPONSE
	// it is up to us (the developers) to decide if that's a problem or not - PANIC
	// deferred functions will still execute
	// Go does not inherently know how to deal with paincs, so if panic is not handled, the program will exit
	if err != nil {
		panic(err.Error())
	}
}
