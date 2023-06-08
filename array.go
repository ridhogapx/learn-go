package main

import ("fmt")

func main() {
	// Declaration
	var fruits = [3]string{"Grapes", "Papaya", "Banana"}
	animals := [3]string{"Cat", "Dog", "Cappybara"}
	// Getting fruit elements
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	fmt.Println(animals)
}