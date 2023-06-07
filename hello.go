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

	// Formating Output
	fmt.Println(a + b)
	fmt.Printf("My name is %v. I'm %v years old \n", MY_NAME, age)
	fmt.Printf("The number is %#v \n", X)
	fmt.Printf("Does he was handsome? %v \n", is_tamvan)
	fmt.Printf("The sound was %v \n", mystr)
	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Second array: %v \n", arr2)
}
