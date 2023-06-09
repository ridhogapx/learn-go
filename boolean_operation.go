package main

import ("fmt")

func main() {
	var nilaiAkhir = 80
	var absensi = 80

	var lulusUjian = nilaiAkhir >= 80
	var lulusAbsensi = absensi >= 80

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)
}