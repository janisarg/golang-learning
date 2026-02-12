package main

import (
	"fmt"
)

func main() {
	// fmt.Println("")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	// if statement without an else.
	if 8%4 ==0 || 7%2 ==0 {
		fmt.Println("Ether 8 or 7 are even")
	}

	// A statement can precede conditionals;
	//  any variables declared in this statement are available 
	// in the current and all subsequent branches.
	//What for, why?
	if num:=9; num<0 {
		fmt.Println(num, "is negative")
	}else if num<10 {
		fmt.Println(num, "has 1 digit")
	}else{
		fmt.Println(num,"has multiple digits")
	}

	// and not outside if statement : error: "undefined:num" 
	// fmt.Println(num)
}