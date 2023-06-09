package main

import ("fmt")

func SayHello() string {
	return "Hello world"
}

func Greet(name string) string {
	message := "Keep going " + name + "-kun!"
	return message
}

func MyFunc(x int, y int) (result int) {
	 result = x * y
	 return
}

func Msg(total int) (result string) {
	result = "I have " + string(total) + "things!"
	return
}

func main() {
	msg := SayHello()
	fmt.Println(msg)

	greet := Greet("Ridho")
	fmt.Println(greet)
	
	multiple := MyFunc(10,5)
	fmt.Println(multiple)

	example := Msg(4)
	fmt.Println(example)
}