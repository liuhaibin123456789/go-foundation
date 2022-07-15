package main

import (
	"log"
	"reflect"
)

type Reindeer string

func (r Reindeer) TakeOff() {
	log.Printf("%q lifts off.", r)
}

func (r Reindeer) Land() {
	log.Printf("%q gently lands.", r)
}

func (r Reindeer) ToggleNose() {
	if r != "rudolph" {
		panic("invalid reindeer operation")
	}
	log.Printf("%q nose changes state.", r)
}

func main() {
	r := Reindeer("rudolph")

	t := reflect.TypeOf(r)

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		log.Printf("%s\n", m.Name)
	}
}
