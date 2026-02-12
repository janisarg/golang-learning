package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5} //declare and initialize an array in one line.
	fmt.Println("dcl:", b)

	b = [...]int{6, 7, 8, 9, 10}
	fmt.Println("dcl:", b)

	// b = [...]int{11, 12: 4, 500}
	// fmt.Println("idx:", b)
	b = [...]int{100, 3: 400, 500} //:, the elements in between will be zeroed.
	fmt.Println("idx:", b)

	var twoD [2][3]int //  multi-dimensional arrays
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{ // initialize multi-dimensional arrays at once too
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println("2d: ", twoD)
}
