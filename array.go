package main

func main() {
	type kata string
	var arra [3]string
	arra[0] = "nama"
	arra[1] = "saya"
	arra[2] = "yayat"
	print(arra[0] + "" + arra[1] + "" + arra[2])
	var arra2 [2]kata
	arra2[0] = "susah"
	arra2[1] = "banget"
	var ara2 = arra2[1]
	print(ara2)
	var arrint [1]int
	arrint[0] = 20
	print(arrint[0])

	var katakata = [2]string{
		"ook", "ok",
	}
	println(katakata[0])
	var angka = [2]int{
		3, 4,
	}
	//println(angka)
	angka[1] = 49
	println(angka[0])
	println(angka[1])
	print(len(angka))
}
