package main

//stuct sama dengan tipe data lainnya yang dapat dimasukkan sebagai parameter
type Costumer struct {
	nama, alamat string
	umur         int
}

func (costumer Costumer) SayHello() {
	println("hello my name is" + costumer.nama)
}

func main() {
	rully := Costumer{nama: "rully"}
	rully.SayHello()
}

// seolah olah struct mempunyai fungsi tapi sebenarnya tidak
// karena fungsi itu berdiri  sendiri
