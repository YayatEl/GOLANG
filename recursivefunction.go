//fungsi yang memanggil dirinya sendiri

package main

func factorial_loop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}
func rekrusif(value int) int {
	if value == 1 {
		return 1
	} else {
		return rekrusif(value-1) * value
	}

}

func main() {
	faktorial := factorial_loop(5)

	println(faktorial)
	println(rekrusif(5))

}
