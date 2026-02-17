package main

import "fmt"

// function intSeq returns another function, which we define anonymously in the body of intSeq.
// The returned function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt)   // reference to "nextInt"  function !? check documentation for more details TODO:
	fmt.Println(nextInt()) //functions result
	fmt.Println(nextInt())
	fmt.Println(nextInt()) // Each call to nextInt increments i by 1 and returns the new value. The variable i is not re-initialized during the calls to nextInt; it retains its value between calls. This is because nextInt is a closure that captures the variable i from its surrounding scope.
	fmt.Println(nextInt()) // state is unique to that particular function, create and test a new one.

	// To confirm that the state is unique to that
	// particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts)
	fmt.Println(newInts())

	fmt.Println(newInts())
	fmt.Println(nextInt())
}
