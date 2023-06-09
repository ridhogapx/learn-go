package main

import ("fmt")

func main() {
	// Go compiler won't yelling if you not used it
	const (
		name string = "Ridho"
		umur int = 23
	)
	
	fmt.Println(name)
	fmt.Println(umur)

}