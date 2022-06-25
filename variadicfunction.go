//varags paramater terakhir dalam function dijadikan varags yang memiliki varags disebut varidic function
package main

func main() {
	jumlah := sumAll(1, 2, 3, 4, 5)
	println(jumlah)
	jumlah2 := sumAll(20, 30)
	println(jumlah2)
	slice := []int{1, 2, 3, 4}

	jumlahslice := sumAll(slice...)
	println(jumlahslice)

	println(withMessage("jumlah ini adalah", 12, 23, 43543))

}
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func withMessage(message string, angka ...int) (string, int) {
	jumlah3 := 0
	for _, v := range angka {
		jumlah3 += v
	}
	return message, jumlah3
}
