package main

import ("fmt")

func main() {
	people := map[string]string{
		"name" : "Ridho",
		"address" : "Jawa Tengah",
	}

	people["hobby"] = "Coding"
	fmt.Println(people)
	fmt.Println(people["name"])
	fmt.Println(people["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Detective Conan"
	book["author"] = "Anon"
	book["genre"] = "Mystery"
	
	fmt.Println(book)
	fmt.Println(book["title"])
	delete(book, "author")
	fmt.Println(book)

}