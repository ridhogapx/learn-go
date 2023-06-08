package main

import ("fmt")

func main() {
	var race string = "Jawir"
	if race == "Jawir" {
		fmt.Println("Aku Jawa!")
	} else {
		fmt.Println("Aku Sunda!")
	}

	month := 3
	
	switch month {
	case 1: 
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4: 
		fmt.Println("April")
	}


	
}