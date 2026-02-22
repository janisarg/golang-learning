package main

import "fmt"

/////////////structure start
type rect struct {
	width, height int
}

//This area method has a receiver type of *rect.
// The method can be called with either a struct or a struct pointer as the receiver.
func (r *rect) area() int {
	return r.width * r.height
}

//Methods can be defined for either pointer or value receiver types.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r *rect) incrSize(a int) int {
	r.width += a
	r.height += a
	return r.width * r.height
}

/////////////structure end

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("width: ", r.width)
	fmt.Println("height: ", r.height)
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())
	fmt.Println("incrSize area: ", r.incrSize(2))
	fmt.Println("width: ", r.width)
	fmt.Println("height: ", r.height)

}
