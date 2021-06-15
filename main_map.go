package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"toffe pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		123456: "mario",
		928394: "luigi",
		482912: "peach",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[123456])

	phonebook[482912] = "balon"
	fmt.Println(phonebook)
}
