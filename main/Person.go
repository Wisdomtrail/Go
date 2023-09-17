package main

import "strconv"

func main() {
	person := new(Person)
	println(person.getPersonInformation())
}

type Person struct {
	Name string
	Age  int
}

func (p Person) getPersonInformation() string {
	p.Age = 45
	p.Name = "sunday"
	return p.Name + " " + strconv.Itoa(p.Age)