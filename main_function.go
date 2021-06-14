package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}
func circleArea(r float64) float64 {
	return math.Phi * r * r
}

func main() {
	sayGreeting("Hello")

	cycleNames([]string{"clouds", "angkes", "mie goreng"}, sayGreeting)

	a1 := circleArea(10.5)

	fmt.Println(a1)
}
