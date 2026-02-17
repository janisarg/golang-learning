package main

import "fmt"

func main() {
	// range on arrays and slices provides both the index and value for each entry.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Printf("sum %d \n", sum)

	//range on arrays and slices provides both the index and value for each entry.
	//Above we didn’t need the index, so we ignored it with the blank identifier _.
	// Sometimes we actually want the indexes though.

	for i, num := range nums {
		if num == 3 {
			fmt.Printf("index %d \n", i)
		}
	}

	//ange on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}
	//range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Printf("key: %s \n", k)
	}

	for _, v := range kvs {
		fmt.Printf("value: %s \n", v)
	}

	//range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and
	// the second the rune itself. See Strings and Runes for more details.

	for i, c := range "go" {
		// fmt.Printf("%d: %c \n", i, c)
		fmt.Println(i, c)
	}

}
