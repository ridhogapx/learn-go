package main

import ("fmt")

func main() {
	// Declaration
	var fruits = [3]string{"Grapes", "Papaya", "Banana"}
	// Getting fruit elements
	for i, fruit := range fruits {
		fmt.Println(fruit)
	}
}