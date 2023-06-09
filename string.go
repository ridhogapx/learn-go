package main

import ("fmt")

func main() {
	mystr := "Good Morning!"
	
	var first_chara byte = mystr[0]

	// Convert from byte to string
	fmt.Printf("First character is: %v \n", string(first_chara))
	fmt.Println(len(mystr))
}