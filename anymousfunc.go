package main

func main() {

	anymousfunc := func(nama string) bool {
		return nama == "anjing"
	}
	registerUser("eko", anymousfunc)
	registerUser("anjing", func(s string) bool {
		return s == "anjing"
	})

}

type blacklist func(string) bool

func registerUser(name string, blblacklist blacklist) {
	if blblacklist(name) {
		println("you're blocked " + name)
	} else {

		println("welcome", name)
	}
}
