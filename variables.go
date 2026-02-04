package main

import "fmt"

func main() {
	// string variable	
	var a = "initial"
	fmt.Println(a)
	//numeric values
	var b, c int = 1,2 
	fmt.Print(b, c ,"\n")
	fmt.Println(b, c)
	fmt.Printf ("b = %d and c = %d\n",b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	//shorthand for declaring and initializing a variable aka var f string = "apple"
	f:= "apple"
	fmt.Println(f)
	//!!! variables declaraed without corresponding initialization are zero-valued 

}