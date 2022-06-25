package main

func main() {
	counter := 0 // variabel berubah
	increament := func() {
		println("increment")
		counter++
	}
	increament()
	increament()
	println(counter)
}
