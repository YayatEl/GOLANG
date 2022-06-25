package main

func main() {
	type angka int
	var angka1 angka = 30
	var angka2 angka = 40
	type boolean bool

	var lulus boolean = 80 < angka1
	var tidaklulus boolean = angka2 < 80
	println(lulus)
	println(tidaklulus)
	var nilaiakhir = tidaklulus && lulus
	println(nilaiakhir)

}
