package main

import "fmt"

func main() {
	// map[keyType]valueType

	// initializing then assigning
	var map1 map[string]int
	map1 = map[string]int{}
	map1["one"] = 1

	// initializing and assigning

	var map2 = map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}

	fmt.Println(map1)
	fmt.Println(map2)

	// map iteration
	var map3 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range map3 {
		fmt.Println(key, "  \t:", val)
	}
}
