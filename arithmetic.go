package main

import ("fmt")

func main() {
	var a int8 = 12
	var b int8 = 5
	var c int8 = a + b

	fmt.Println(c)

	// Augmented assignment
	var i int8 = 5
	i += 8

	fmt.Println(i)
}