package main

import (
	"fmt"
	"slices"
)

func main() {
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	// An uninitialized slice equals to nil and has length 0.
	var s []string //array of strings, but the length is not defined yet.
	fmt.Println("uninit", s, s == nil, len(s) == 0)
	// fmt.Println("")
	// To create a slice with non-zero length, use the builtin make.

	// By default a new slice’s capacity is equal to its length;
	//  if we know the slice is going to grow ahead of time,
	// it’s possible to pass a capacity explicitly as an additional parameter to make.

	s = make([]string, 3) //  make a slice of strings of length 3 (initially zero-valued).
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	///add new value in slice
	s = append(s, "d") //append new value in slice, and assign it to s again.
	s = append(s, "e", "f")
	fmt.Println("append:", s)
	fmt.Println("get:", s[len(s)-1])
	// create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	d := make([]string, 0)
	copy(d, s)
	fmt.Println("cpy empty:", d)
	fmt.Println("get 2:", s)

	//Slices support a “slice” operator with the syntax slice[low:high].
	// For example, this gets a slice of the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("sl1:", l)
	//This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("sl2: ", l, " len:", len(s))
	//slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("sl3: ", l, " len:", len(s), " s= ", s)
	//declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	//The slices package contains a number of useful utility functions for slices.

	// t2 := []string{"h", "g", "i"} //not eq
	t2 := []string{"h", "g", "i"} //eq
	if slices.Equal(t, t2) {
		fmt.Println("equal:", t, t2)
	} else {
		fmt.Println("not equal:", t, t2)
	}

	twoD := make([][]int, 3) //create a 2D slice with 3 rows.
	//for i := range twoD {
	for i := range 3 {
		innerLen := i + 1
		// twoD[i] = make([]int, 4) //create 4 columns for each row.
		twoD[i] = make([]int, innerLen) //create different columns for each row.
		for j := range twoD[i] {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
