package main

import (
	"fmt"
	"maps"
	// "maps"
)

func main() {
	// Create a map to store the ages of people
	m := make(map[string]int)
	m["k1"] = 7 //set key/value
	m["k2"] = 13
	fmt.Println("map: ", m)

	v1 := m["k1"]
	// v1 = m["k1"]
	fmt.Println("v1: ", v1)

	v3 := m["k1"]
	fmt.Println("v3: ", v3)

	// m["k1"] = 99
	// fmt.Println("v3: ", v3)

	fmt.Println("len: ", len(m))

	delete(m, "k2") //The builtin delete removes key/value pairs from a map.
	fmt.Println("map: ", m)

	clear(m) //To remove all key/value pairs from a map, use the clear builtin.
	fmt.Println("map: ", m)

	// optional second return value when getting a value from a map indicates
	// if the key was present in the map. This can be used to disambiguate
	// between missing keys and keys with zero values like 0 or "".
	// Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// maps package contains a number of useful utility functions for maps.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
