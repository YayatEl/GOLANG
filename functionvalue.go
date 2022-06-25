package main

/*
function di golang dapat berupa tipe data dan dapat disimpan dalam variabel function digolang merupakan first class citizen
*/
func main() {
	sayGoodbye := getGoodbye
	println(sayGoodbye("john"))
}

func getGoodbye(name string) string {
	return "goodbye " + name
}
