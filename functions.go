package main

import (
	"fmt"
)

// simple functions
func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// /multiple return values
func vals() (int, int) {
	return 3, 7
}

//fignja
// func lja() (...int) {
// 	return 1, 2, 3, 4, 5
// }

// Variadic Functions
func mySum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)

	a, b := vals()
	fmt.Println("vals(a)", a)
	fmt.Println("vals(b)", b)
	//If you only want a subset of the returned values, use the blank identifier _.
	_, c := vals()
	fmt.Println(c)

	mySum(1, 2)
	mySum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	mySum(nums...)

	// opa:=lja()
	// fmt.Println("opa:",opa)

}
