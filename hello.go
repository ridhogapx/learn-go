package main
import ("fmt")

func main() {
	 // var str string = "Yoooo it's me, dawg!"

	// Grouped var
	var (
		 a int = 15
		 b int = 20
	)

	// Grouped Constant 
	const (
		MY_NAME string = "Ridho"
		X int = 23
		is_tamvan bool = true
	)
	// Formating Output
	fmt.Println(a + b)
	fmt.Printf("My name is %v\n", MY_NAME)
	fmt.Printf("The number is %#v\n", X)
	fmt.Printf("Does he was handsome? %v", is_tamvan)
}