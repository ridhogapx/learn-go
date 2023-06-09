package main

import ("fmt")

func main() {
	// Number 
	var age int = 17
	fmt.Println(age)

	// String
	var name string = "Ridho"
	fmt.Printf("My name is %v \n", name)

	// Untyped
	greet := "Welcome!"
	fmt.Printf("He's said: %v \n", greet)

	var (
		pet string = "Cat"
		pet_name string = "Momosuke"
		pet_age int = 1
	)

	fmt.Printf("I have a %v. Its name is %v. He's %v years old \n", pet, pet_name, pet_age)

	var Person string

	Person = "John"
	fmt.Println(Person)

	Person = "Doe"
	fmt.Println(Person)
	
	Person = "Jane"
	fmt.Println(Person)

	var cat = "Momo"
	fmt.Println(cat)

	var cat_age int8 = 1
	fmt.Println(cat_age)
	
	// Tanpa keyword var
	race := "Jawir"
	fmt.Println(race)

	var (
		student_1 = "Budi"
		student_2 = "Andy"
	)

	fmt.Println(student_1)
	fmt.Println(student_2)

}