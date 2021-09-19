package golang

import "fmt"

func fun() {
	// Variables are declared with:
	// var name type
	var foo string = "Foobar!"
	// Go has print functions similar to C.
	fmt.Println(foo)
}

// Similar to variables, a functions return type is set after the name and parameters.
func summation(n int) int {
	// var allows the programmer to explicitly declare the type.
	var sum int = 0
	// := (the walrus operator) sets the type to whatever is being assigned.
	for i := 0; i < n; i++ {
		sum++
	}
	return sum
}

// You can even return multiple types with functions.
func what() (int, string) {
	return 42, "answer to life the universe and everything"
}

// Any data member that starts with a capitol letter is visible when a package is imported.
func Samples() {
	fun()

	s := summation(5)
	fmt.Println(s)

	// How to assign variables with a multi-type return function.
	answer, question := what()
	fmt.Println(question, "is", answer)
}
