// interface adalah tipe data yang abstrak tidak memiliki implementasi langsung
// interface berisikan definisi-definisi method
// biasanya interface digunakan sebagai kontrak
package main

type HasName interface {
	GetName() string
}

func (person Person) GetName() string {
	return person.name
}
func (animal Animal) GetName() string {
	return animal.name
}
func sayHI(hasname HasName) {
	println("hello", hasname.GetName())
}

type Person struct {
	name string
}
type Animal struct {
	name string
}

func main() {
	var eko Person
	var Kuda Animal
	eko.name = "eko"
	Kuda.name = "kuda"
	kucing := Animal{name: "kucing"}
	sayHI(eko)
	sayHI(Kuda)
	sayHI(kucing)
}
