package main

func main() {
	person := map[string]string{

		"name":    "eko",
		"address": "antang",
	}
	println(person["name"])
	println(person["address"])
	//person["title"] = "biodata"
	println(person)
	println(person["title"])
	delete(person, "title")
}
