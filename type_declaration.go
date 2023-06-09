package main

import ("fmt")

func main() {
	type isGraduated bool

	var name string = "Ridho"
	var status isGraduated = true

	fmt.Printf("My name is: %v \n", name)
	fmt.Println(status)
}