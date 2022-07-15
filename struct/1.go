package main

import "fmt"

type Name struct {
	FirstName string
	LastName  string
}

type Person struct {
	Name
	Age int
}

func main() {
	me := new(Person)
	me.Name.FirstName = "Fluoride"
	me.Name.LastName = "Hydrogen"
	fmt.Printf("%#v\n", me)
	//me: main.Person{Name:main.Name{FirstName:"Fluoride", LastName:"Hydrogen"}, Age:0}

	//me1 := new(Person)
	//me1.FirstName = "Fluoride"
	//me1.LastName = "Hydrogen"
	//fmt.Println(me1)
	// me := &Person{FirstName: "Fluoride", LastName: "Hydrogen", Age: 1}
}
