package main

import ("fmt")

func main() {
	var a int8 = 10
	var b int8 = 10

	var c int8 = 12
	var d int8 = 50

	var result = a == b 
	fmt.Println(result)

	fmt.Println(c > d )
	fmt.Println( d > c)
}