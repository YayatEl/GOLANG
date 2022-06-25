package main

//slice mirip array yang bisa berubah ukurannya
//slice memliki 3 data pointer length capacity
func main() {
	var bulan = [...]string{
		"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember",
	}
	var slice1 = bulan[1:4]
	var capslice = cap(slice1)
	println(slice1)
	//len panjang, cap kapasitas append make copy
	println(slice1[0])
	slice1[1] = "meilagi"
	println(slice1[1])
	println(len(slice1))
	println(capslice)
	names := [...]string{
		"adi", "dudung", "budi",
	}
	println(names[0:])

	//	array2 := [...]string{

	//		"mangga", "jeruk", "pisang",
	//	}
	slice2 := bulan[3:]
	slice3 := append(slice2, "okay")
	println(slice3[8])

	slice4 := make([]string, 2, 5)
	slice4[0] = "eko"
	slice4[1] = "budi"
	println(slice4[1])
	copyslice := make([]string, len(slice4), cap(slice4))
	copy(copyslice, slice4)
	println(copyslice[0])
	iniarray := [...]int{1, 2, 3, 4, 5}
	inslice := []int{1, 2, 3, 4, 5}
	println(iniarray[1])
	println(inslice[1])
}
