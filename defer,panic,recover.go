package main

func main() {
	runApplication(0)
}
func logging() {
	println("selesai memanggil function")
}
func runApplication(value int) {
	defer logging()
	result := 0
	result = 10 / value
	println(result)
}
