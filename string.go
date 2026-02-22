package main

import (
	"fmt"
	"unicode/utf8"
)

// A string is a slice of bytes. Strings are immutable: once created, it is not possible to change the contents of a string.
// The range of a string iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself.
func main() {
	const s = "Hallo  āēšķļž"
	fmt.Printf("%s \n", s)
	fmt.Printf("length in bytes: %d \n", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
		// fmt.Printf("%d: %x : %d : %c\n", i, s[i], s[i], s[i]) // %x prints the byte in hexadecimal, %d prints the byte in decimal and %c prints the byte as a character.
	}
	fmt.Println()
	fmt.Printf("length in runes(chars): %d \n", utf8.RuneCountInString(s))

	//A range loop handles strings specially and decodes each rune along with its offset in the string.
	for idx, runeValu := range s {
		fmt.Printf("%#U starts at %d \n", runeValu, idx)
	}
	//using the utf8.DecodeRuneInString function explicitly.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		// This demonstrates passing a rune value to a function.

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	// Values enclosed in single quotes are rune literals. We can compare a rune value to a rune literal directly.

	if r == 'a' {
		fmt.Println("found a")
	} else if r == 'ā' {
		fmt.Println("found aaaa")
	}
}
