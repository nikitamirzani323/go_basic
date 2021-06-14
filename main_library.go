package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greetings := "hello there friends"

	fmt.Println(strings.Contains(greetings, "hello"))
	fmt.Println(strings.ReplaceAll(greetings, "hello", "hay"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, "th"))
	fmt.Println(strings.Split(greetings, " "))

	fmt.Println("Original string value : ", greetings)

	ages := []int{45, 20, 25, 30, 75, 89}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 90)
	fmt.Println(index)

	names := []string{"adam", "tary", "shinta", "lucas", "edo"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "tary"))
	fmt.Println(sort.SearchStrings(names, "batara"))
}
