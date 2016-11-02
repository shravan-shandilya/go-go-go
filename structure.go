package main

import "fmt"

type Person struct{
	name string
	age int
	*Person
}


func main(){
	p := Person{name:"first_guy",age:99,Person:nil}
	q := Person{"second_guy",98,nil}
	r := &Person{"pointer_guy",23,nil}


	fmt.Println("First_guy:",p)
	fmt.Println("Second_guy:",q)
	fmt.Println("third_guy accessed via pointer",*r)

	fmt.Println("accessing attributes")
	fmt.Println("Pointer to third guy",r)

	fmt.Println("third guy's name",r.name)
	fmt.Println("third guy's age",r.age)
	fmt.Println("***************************")
	fmt.Println("No matter how you are acessing atributes of a struct, you use <pointer>.<attribute> or <instance>.<attribute>")
	fmt.Println("***************************")

	s := Person{}
	fmt.Println(s)
	s.name = "satoshi nakamoto"
	s.age = 999
	s.Person = &s

	fmt.Println(s)

	fmt.Println("associating method with structure")
	s.printInfo()
	s.ageing(10)
	s.printInfo()
	s.ageing(-100)


}

func (p *Person) printInfo() {
	fmt.Println("****************************")
	fmt.Println("Name:",p.name)
	fmt.Println("Age:",p.age)
	fmt.Println("Father:",p.Person.name)
	fmt.Println("****************************")
}

func (p *Person) ageing(age int) int{
	p.age = p.age + (age)
	return p.age
}
