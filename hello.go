package main

import (
	"fmt"
)

func main() {
	// var str string = "Yoooo it's me, dawg!"

	// Grouped var
	var (
		a int = 15
		b int = 20
	)

	// Untyped
	mystr := "Whooosh"

	// Grouped Constant
	const (
		MY_NAME   string = "Ridho"
		age       int    = 17
		X         int    = 23
		is_tamvan bool   = true
	)

	// Array
	arr := [3]int{23, 1, 2006}
	var arr2 = [...]string{"Foo", "Bar"}
	fruits := [5]string{"Mango", "Orange", "Pinapple"}

	arr2[1] = "Whooosh"

	// Slices
	slice_a := []int{5, 10, 20, 15, 7, 10}
	slice_b := []string{"Orange", "Banana", "Apple"}

	// Formating Output
	fmt.Println(a + b)
	fmt.Printf("My name is %v. I'm %v years old \n", MY_NAME, age)
	fmt.Printf("The number is %#v \n", X)
	fmt.Printf("Does he was handsome? %v \n", is_tamvan)
	fmt.Printf("The sound was %v \n", mystr)
	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Second element from 2nd array is: %v \n", arr2)
	fmt.Printf("List of fruits: %v \n", len(fruits))
	fmt.Printf("Slice A = %v \n", slice_a)
	fmt.Printf("Slice B = %v \n", slice_b)
	fmt.Printf("The length of Slice A: %v \n", len(slice_a))
	fmt.Printf("The length of Slice B: %v \n", len(slice_b))

}
