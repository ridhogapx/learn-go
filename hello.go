package main

import (
	"fmt"
)

func greet(name string) string {
	msg := "Goodluck " + name
	return msg
}

func sum(a int, b int) int {
	return a + b
}

func recursion_example(num int) int {
	if num == 15 {
		return 0
	}
	fmt.Println(num)
	return recursion_example(num + 1)
}

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
	arr := [10]int{23, 1, 2006, 50, 11, 43, 14, 90, 24, 54}
	var arr2 = [...]string{"Foo", "Bar"}
	fruits := [5]string{"Mango", "Orange", "Pinapple"}

	arr2[1] = "Whooosh"

	// Slices
	slice_a := []int{5, 10, 20, 15, 7, 10, 15, 16, 17}
	slice_b := []string{"Orange", "Banana", "Apple"}
	deretan := []int{23, 20, 25}
	gabungan := append(slice_a, deretan...)

	// The syntax is slice_name := array[index_start:index_end]
	slice_from_array := arr[3:8]
	// Creating slices with make() function
	slice_with_make := make([]int, 5, 10)
	// Appending / Adding elements to slice
	slice_c := append(slice_a, 69, 233, 59, 100)
	more_fruits := append(slice_b, "Papaya", "Melon", "Grapes", "Jackfruit")
	// Coba-coba
	slices_fruit := slice_b[1:3]

	// Memory Efficiency
	animals := []string{"Cat", "Dog", "Girrafe", "Lion", "Tiger", "Chicken"}
	needed := animals[:5]
	animalsCopy := make([]string, len(needed))
	copy(animalsCopy, animals)

	// Conditional
	num_a := 15
	num_b := 20

	if num_a > num_b {
		fmt.Printf("A lebih besar daripada B \n")

	} else {
		fmt.Printf("B lebih besar daripada A \n")
	}

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
	fmt.Printf("Slice from variable 'arr': %v \n", slice_from_array)
	fmt.Printf("Slice with make function: %v \n", slice_with_make)
	fmt.Printf("Slice C: %v \n", slice_c)
	fmt.Printf("More fruits: %v \n", more_fruits)
	fmt.Printf("Gabungan dari slices: %v \n", gabungan)
	fmt.Printf("Slices Fruit: %v \n", slices_fruit)
	fmt.Printf("Copy of %v \n", animalsCopy)

	// If else
	x := 5
	y := 10

	if x >= 5 {
		fmt.Printf("X lebih besar daripada %v \n", x)
		if y >= 10 {
			fmt.Printf("Y lebih besar daripada %v \n", y)
		}
	}

	// Switch
	month := 3

	switch month {
	case 1:
		fmt.Printf("Januari \n")
	case 2:
		fmt.Printf("Februari \n")
	case 3:
		fmt.Printf("Maret \n")
	}

	// For loop
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("---------------------------")

	// With break keyword
	for j := 0; j < 10; j++ {
		if j == 8 {
			break
		}
		fmt.Println(j)
	}

	// Get element array
	veggie := [3]string{"Spinach", "Tomato", "Carrot"}
	for el := 0; el < len(veggie); el++ {
		fmt.Println(veggie[el])
	}

	fmt.Println("---------------------------")
	for i, val := range veggie {
		fmt.Println(i, val)
	}

	fmt.Println("---------------------------")

	for _, val := range veggie {
		fmt.Println(val)
	}

	fmt.Println("---------------------------")

	// Function
	fmt.Println(greet("Andy"))
	penjumlahan := sum(34, 70)
	fmt.Println(penjumlahan)
	fmt.Println("---------------------------")
	// Recursion
	recursion_example(3)
	fmt.Println("---------------------------")
	// Struct
	type Somebody struct {
		name string
		age  int
		job  string
	}

	var person Somebody
	person.name = "Galih"
	person.age = 17
	person.job = "Programmer"

	fmt.Printf("Name: %v \n", person.name)
	fmt.Printf("Age: %v \n", person.age)
	fmt.Printf("Job: %v \n", person.job)
	fmt.Println("---------------------------")
	// Map
	// Syntax: map[KeyType]ValueType{key1: value1}
	weapon := map[string]string{"name": "Vandal", "type": "Rifle"}
	fmt.Printf("Weapon Map: %v \n", weapon)
	// Map with make()
	area := make(map[string]string) // The map is empty now
	area["name"] = "Lotus"
	area["Location"] = "China"
	area["release"] = "1 October"

	// Deleting map element
	delete(area, "release")
	delete(area, "name")
	fmt.Println(area)

	// Updating map
	weapon["price"] = "$ 2000"

	fmt.Printf("Weapon: %v \n", weapon["name"])
	fmt.Printf("Price: %v \n", weapon["price"])
	fmt.Printf("Area: %v \n", area["name"])
}
