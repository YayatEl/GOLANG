package main

/*multiline
 */

func endprogram() {
	println("tes panic fungsi")
	messege := recover()
	println("aplikasi kembali berjalan", messege)
}
func runApp(error bool) {
	defer endprogram()
	if error {
		panic("error")
	}
	messege := recover()
	println("aplikasi kembali berjalan", messege)
}
func main() {
	runApp(true) //komentar single line
}
