package main

func main() {

	println(returnNamedValues())
	namadepan, namatengah, namabelakang := returnNamedValues()
	println(namadepan)
	println(namatengah)
	println(namabelakang)
}
func returnNamedValues() (firstNama, middleName, lastName string) {
	firstNama = "marco"
	middleName = "The"
	lastName = "phoenix"
	return
}
