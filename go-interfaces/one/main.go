// INTERFACES DEFINE BEHAVIOUR (NOT DATA) hence interfaces wrap functions unlike structs
package main

import (
	"fmt"
	"os"
)

type writer interface {
	write([]byte) (int, error)
}

type consoleWriter struct{}

// you don't have to use interfaces with structs, it can be any kind of CUSTOM type
type intCounter int

type incrementer interface {
	increment() int
}

func (ic *intCounter) increment() int {
	*ic++
	return int(*ic)
}

func (cw consoleWriter) write(data []byte) (int, error) {
	/*
				fleshing out Println - more on interfaces
				signature: func Println(a ...any) (n int, err error) { return Fprintln(os.Stdout, a...) }
				signature: func Fprintln(w Writer, a ...any) (n int, err error)
				any is an alias for the empty interface
				therefore:
				fmt.Println(string(data)) == fmt.Fprintln(os.Stdout, string(data))

				and os.Stdout is a VARIABLE
				in its definition, it implements the Writer interface - has a function NewFile() that returns a *File (pointer to a File object)
				os.Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout") *File
				uintptr(syscall.Stdout) is a File Descriptor for standard output (==1) - system calls and processes, etc

				therefore, os.Stdout is of type *File (returned by NewFile) - an open File pointing to standard output
				and variables of type *File have a method Write:
				func (f *File) Write(b []byte) (int, error) {...}
				and the Writer interface wraps the basic write function
				type Writer interface {
					write(p []byte) (int, error)
				}
		 		therefore, os.Stdout implements the Writer interface - is also of type Writer - IMPLICIT IMPLEMENTATION
	*/
	// func Fprintln(w Writer, a ...any) (n int, err error)
	n, err := fmt.Fprintln(os.Stdout, string(data))
	return n, err
}

func main() {
	// interfaces in action - polymorphism
	var w writer = consoleWriter{}

	// does not know what it is writing to, handled elsewhere
	w.write([]byte("Hello.")) // type conversion - string to []byte

	// some more interface action

	myInt := intCounter(0)       // intCounter object
	var inc incrementer = &myInt // pointer to an intCounter (peep method definition)

	for i := 0; i < 10; i++ {
		fmt.Println(inc.increment())
	}
}
