package main

import "fmt"

func main() {
	age := 35
	name := "Shain"
	fmt.Print("Hello, ")
	fmt.Print("world \n")
	fmt.Print("new line \n")

	fmt.Println("hello ninja")
	fmt.Println("goodbay ninja")
	fmt.Println("my age is ", age, " and my name is ", name)

	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f points \n", 225.55)
	fmt.Printf("you scored %0.2f points \n", 225.55)
	fmt.Printf("you scored %0.1f points \n", 225.55)

	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
}
