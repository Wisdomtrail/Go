package main

func main() {
	person := new(Person)
	println(person.getPersonInformation())
}

type Person struct {
	Name string
	Age  int
}

func (p Person) getPersonInformation() interface{} {

	return nil
}
