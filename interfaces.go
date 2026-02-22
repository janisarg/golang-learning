package main

//Interfaces are named collections of method signatures.
//https://www.youtube.com/watch?v=SX1gT5A9H-U

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Measurable interface {
	perim() (float64, error)
}

// this is a "private" interface, since it starts with a lowercase letter.
// It defines a geometry interface with two methods: area and perim.
type geometry interface {
	//  interfaces can embed other interfaces, just like structs can embed other structs.
	Shape
	Measurable
	// area() float64
	// perim() float64
	//	opa() float64  // this method is not implemented by the rect and circle types, so they do not satisfy the geometry interface.
}

type CalulationError struct {
	msg string
}

// The error interface is a built-in interface that has a single method, Error, which returns a string.
// one of the most common uses of interfaces is to implement the error interface, which allows us to create custom error types.
func (ce CalulationError) Error() string {
	return ce.msg
}

// For our example we’ll implement this interface on rect and circle types.
// rect is a private struct with two fields, width and height, of type float64.
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// / We implement the geometry interface on rect by implementing the area and perim methods on it.
// some of the methods have a value receiver and some have a pointer receiver.
// This is to show that it doesn’t matter which you use - the interface will work with either.
// very confusing comparing to the Java and C# interfaces, where you have to implement all the methods with the same receiver type.
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() (float64, error) {
	return 2*r.width + 2*r.height, nil
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() (float64, error) {
	if c.radius < 0 {
		return 0, CalulationError{msg: "negative radius"}
	}
	return 2 * math.Pi * c.radius, nil
}

// The measure function takes a geometry interface value and calls its area and perim methods.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func descShape(ge geometry) {
	fmt.Println(ge)
	fmt.Println(ge.area())
	fmt.Println(ge.perim())
	// fmt.Println(ge.opa())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
	fmt.Println("rect: ", r)
	fmt.Println("circle: ", c)
	//fmt.Println("circle: ", c.measure()) // this will not compile, because the circle type does not have a measure method, even though it satisfies the geometry interface.
	i := geometry(r) // this is a type assertion, it converts the rect value to a geometry interface value.
	fmt.Println("geometry: ", i)
	mysteryB := interface{}(10) // this is an empty interface, it can hold any value, but it does not have any methods, so we cannot call any methods on it.
	descValue(r)
	descValue(mysteryB)
	retrievedInt, ok := mysteryB.(int)
	// this is a type assertion, it tries to convert the mysteryB value to an int,
	// and returns the int value and a boolean indicating whether the conversion was successful.
	if ok {
		fmt.Println("mysteryB is an int:", retrievedInt)
	} else {
		fmt.Println("mysteryB is not an int")
	}

	descShape(r)
	descShape(c)
	descShape(i)
	// descShape(mysteryB) // this will not compile, because the mysteryB value does not satisfy the geometry interface,
	// even though it is an int, which is a basic type that does not have any methods.
	//fmt.Println("geometry: ", i.measure()) // this will not compile, because the geometry interface does not have a measure method,
	// even though the rect type satisfies the geometry interface.)

}

func descValue(t interface{}) {
	fmt.Printf("value: %v, type: %T\n", t, t)
}
