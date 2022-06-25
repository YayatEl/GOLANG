package main

func main() {
	for i := 0; i < 10; i++ {
		println("perulangan ke:", i)

		if i == 5 {
			break
		}
	}
	//continue akan menghentingan perulangan dan melanjutkan ke post statment
	for i := 0; i < 10; i++ {
		//println("perulangan ke: ", i)
		if i%2 == 0 {
			continue
			//println("perulangan ke: ", i)
		}
		println("perulangan ke: ", i)
	}
}
