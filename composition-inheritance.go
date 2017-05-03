package main

import (
	"fmt"
)

type Pet struct {
	name string
}

type Dog struct {
	*Pet
	Breed string
}

func (p *Pet) EatFood() string {
	return "chomp chomp chomp!"
}

func (p *Pet) Speak() string {
	return fmt.Sprintf("my name is %v", p.name)
}

func (d *Dog) Speak() string {
	return fmt.Sprintf("%v and I am a %v", d.Pet.Speak(), d.Breed)
}

func (d *Dog) Bark() string {
	return "Woof Woof"
}

func main() {
	d := Dog{Pet: &Pet{name: "satoshi"}, Breed: "beagle"}

	//method get overrided by dog's method
	fmt.Println(d.Speak())

	//method which is not in Pet,but offered by Dog
	fmt.Println(d.Bark())

	//method which is not in Dog, but offered by all Pet
	fmt.Println(d.EatFood())
}
