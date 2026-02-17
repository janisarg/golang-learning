package main

import "fmt"

//We’ll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr.
// zeroval has an int parameter, so arguments will be passed to it by value.
// zeroval will get a copy of ival distinct from the one in the calling function.
func zeroval(ival int) {
	ival = 0
}

//zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
// The *iptr code in the function body then dereferences the pointer from its
// memory address to the current value at that address. Assigning a value to a
//  dereferenced pointer changes the value at the referenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	// A pointer holds the memory address of a value. The type *T is a pointer to a T value. Its zero value is nil.
	var p *int
	fmt.Printf("the value of p is %v \n", p)
	// The & operator generates a pointer to its operand.
	i := 42
	p = &i
	fmt.Printf("the value of p is %v \n", p)
	fmt.Printf("the value of *p is %v \n", *p)

	i = 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)

}
