package main

import ("fmt")

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	slice_a := months[4:10]
	fmt.Println(slice_a)
	fmt.Printf("Length of Slice A: %v \n", len(slice_a))
	fmt.Printf("Capacity of Slice A: %v \n", cap(slice_a))
	slice_a[1] = "Juni lagi"
	fmt.Println(months)

	slice_b := months[0:3]
	fmt.Println(slice_b)

	slice_c := append(slice_b, "Miawww")
	fmt.Println(slice_c)
	fmt.Println(months)

	newSlice := make([]int8, 3, 6)
	newSlice[0] = 23
	newSlice[1] = 1
	newSlice[2] = 123

	fmt.Println(newSlice)

	copySlice := make([]int8, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// IMPORTANT !!!
	ini_array := [...]string{
		"Foo",
		"Bar",
	}
	ini_slice := []string{
		"Foo",
		"Bar",
	}

}