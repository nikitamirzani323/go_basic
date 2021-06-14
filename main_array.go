package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{20,35,35}
	var ages = [3]int{20, 35, 35}

	names := [4]string{"andi", "mario", "lukas", "setan"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	var scores = []int{100, 25, 10}
	scores[2] = 30
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)
}
