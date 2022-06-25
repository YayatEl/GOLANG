package main

func main() {
	a := 1
	nama := "eko"
	//b := 20

	switch a {
	case 10:
		println("budi")
		a = 20
	case 20:
		println("ok")
	default:
		println("okok")
		break

	}
	switch length := len(nama); length < 1 {
	case true:
		println(a)
	default:
		println("salah")

	}
	length := len("yayat")
	switch {
	case length > 5:
		println(length)
	case length < 5:
		println("besar bro")
	default:
		break

	}
}
