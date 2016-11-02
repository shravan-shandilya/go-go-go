package main

import "fmt"

func main(){
	some_integer := 2342
	pointer_to_an_integer := &some_integer
	fmt.Println("Value accessed via its pointer:",*pointer_to_an_integer)
	fmt.Println("Value stored in pointer variable:",pointer_to_an_integer)

	fmt.Println("Some pointer arithmatic")
	pointer_to_an_integer = (*int)((int)pointer_to_an_integer + 1)
	fmt.Println("after adding 1 :",pointer_to_an_integer)
	fmt.Println("accessing now:",*pointer_to_an_integer)


	pointer_to_an_integer = nil
	fmt.Println("Trying a null derefernce")
	fmt.Println(*pointer_to_an_integer)
}
