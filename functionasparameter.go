package main

func main() {
	sayHelloWithFilter("anjing", spamFilter)
	funcvariable := spamFilter

	sayHelloWithFilter("anjing", funcvariable)
}

type Filtered func(string) string

func sayHelloWithFilter(name string, filter Filtered) {
	nameFiltered := filter(name)
	println("hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "jangan bicara kotor"
	} else {

		return name
	}
}
