package main

func main() {
	for i := 0; i < 10; i++ {
		println("perulangan ke", i)

	}
	count := 1
	for count < 10 {
		println("ulang")
		count++
	}
	slice := []string{"eko", "budi", "yayat", "goblok", "deri"}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	for _, value := range slice {
		println(value)
	}

	mapping := make(map[int]string)
	mapping[1] = "eko"
	mapping[2] = "budi"
	for key, v := range mapping {
		println(key, v)
	}

}
