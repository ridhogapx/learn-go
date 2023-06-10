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

func multiVal() (string, string) {
	return "Hello", "World"
}

// Variadic
func exVariadic(a ...int) int {
	num := 0
	for _, el := range a {
		num += el
	}
	return num
}

// Anonymous
type Coba func(string) string
func Cb(foo string,anon Coba)  {
	anon(foo)
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

	_, lastWord :=  multiVal()
	fmt.Println(lastWord)

	exNumber := []int{5,10,20}
	fmt.Println(exVariadic(2,5,1,6))
	fmt.Println(exVariadic(exNumber...))

	tanpa_nama := Cb("Hello", func(name string) string {
		return name == "Bar"
	})

}