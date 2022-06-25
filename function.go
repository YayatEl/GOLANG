package main

func main() {
	print()
	namabelakang := "mambo"
	menambahkan("budi", namabelakang)
	hasilrern := sretunVal("")
	println(hasilrern)
	println(returnmulti())

	va1, va2, va3 := returnmulti()

	println(va1, va2, va3)
	va4, va5, _ := returnmulti()
	println(va4, va5)
}

//parameter is value outside function
func print() {
	println("memprint print")
}

//function with parameter
func menambahkan(firstname string, lastname string) {
	println("nama saya ", firstname, lastname)
	println("nama saya ", lastname)
}

//function return value

func sretunVal(nm string) string {
	nameplus := nm + "mama"

	if nm == "" {
		return "kosong"
	} else {

		return nameplus
	}
}

//returning multiple values in function
func returnmulti() (string, string, int) {
	name1 := "adi"
	name2 := "dodi"
	angka := 10
	return name1, name2, angka
}
