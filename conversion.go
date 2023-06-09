package main

import ("fmt") 

func main() {
	var age int8 = 23
	fmt.Println(age)
	
	// Conversion type data
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai)
}