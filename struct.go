package main

/**

sturct adalah template untuk mengbaunggkan 0 atau lebih tipe data lainnya dalam kesatuan
sturct biasanya representasi dari data yang kita buat
data struct disimpan dalam field
sturct adalah kumpulan data filed secara sederhana
*/

type OrangBiasa struct {
	name, Addres string
	umur, gaji   int
}

func main() {
	var mahasiswa OrangBiasa
	mahasiswa.gaji = 1000
	mahasiswa.name = "eko"
	mahasiswa.Addres = "antang"
	mahasiswa.umur = 20
	println(mahasiswa.gaji)
	println(mahasiswa.name)
	println(mahasiswa.Addres)
	println(mahasiswa.umur)

	//cara lain membuat sturct
	mahasiswa2 := OrangBiasa{
		name: "joko",
		umur: 20,
	}
	println(mahasiswa2.name)

	mahasiswa3 := OrangBiasa{"widi", "antang", 20, 20000} //harus sama jumlah field
	println(mahasiswa3.name)

}
