package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Winaldo Satryadi Manurung")

	// string length
	fmt.Println(len("Winaldo"))

	// print ASCII dec at index [i]
	fmt.Println("Winaldo"[0])
	fmt.Println("Winaldo Satryadi Manurung"[1])

	// checking strings contains another string
	fmt.Println(strings.Contains("Winaldo Satryadi", "Manurung"))
	fmt.Println(strings.Contains("Winaldo Satryadi", "Satryadi"))

	// splitting string into slice
	fmt.Println(strings.Split("Winaldo Satryadi Manurung", " "))

	// string methods
	fmt.Println(strings.ToLower("Winaldo Satryadi Manurung"))
	fmt.Println(strings.ToUpper("Winaldo Satryadi Manurung"))
	fmt.Println(strings.ToTitle("Winaldo Satryadi Manurung"))

	// trim whitespace
	fmt.Println(strings.Trim("      Winaldo Satryadi     ", " "))

	// find and replace
	fmt.Println(strings.ReplaceAll("Winaldo Manurung", "Manurung", "Oke"))
}
