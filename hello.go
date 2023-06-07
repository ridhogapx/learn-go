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
	)
	// Formating Output
	fmt.Println(a + b)
	fmt.Printf("My name is %v", MY_NAME)
	fmt.Printf("The number is %#v", X)
}