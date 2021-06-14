package main

import "fmt"

func main() {
	//String
	var nameOne string = "Mario"
	var nametwo string = "Luigi"
	var namethree string

	fmt.Println(nameOne, nametwo, namethree)

	nameOne = "peach"
	namethree = "browser"
	fmt.Println(nameOne, nametwo, namethree)

	namefour := "yoshi"
	fmt.Println(namefour)

	//int
	var ageOne int = 20
	var ageTwo int = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint = 25
	fmt.Println(numOne, numTwo, numThree)
}
