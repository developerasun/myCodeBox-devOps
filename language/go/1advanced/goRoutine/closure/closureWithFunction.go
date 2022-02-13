package main
// Any go code file starts by defining the package
// package name of a file should be the name of a folder
// that the file blongs to : to simply put - package name == folder name
// if package name == "main", then it becomes an entry point of your program.

import (
	"fmt"
)

// Create a function that takes a function as argument
func AddMul(value int, myFunc func(a, b int) int) int {
	return value * myFunc(value, value)
}

// Create a closure environment in golang
func incrementer() func() int {
	// Initializes i
	var i int = 5

	// fuction incrementer returns func()
	return func() int {
		i += 2
		return i
	}
}

func main() {

	// Declare a function and store it in a variable : function value
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println("function variable: ", add(5, 7))
	fmt.Println("function args: ", AddMul(3, add))

	// Declare a variable with the incrementer function
	LetsIncrease := incrementer()

	// Create arithmetic progression : a1 = 5, d = 2
	for i := 0; i < 5; i++ {
		save := LetsIncrease()
		fmt.Println("term: ", save)
	}

	ClosureWithGoRoutine()
}
