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

	// Manual declaration
	var names [3]string
	names[0] = "Ridho"
	names[1] = "Galih"
	names[2] = "Pambudi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Other declaration
	var myNumber = [3]int8{
		50,
		75,
		100,
	}

	fmt.Println(myNumber)
	fmt.Printf("Length of names: %v", len(names))
}