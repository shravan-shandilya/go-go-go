package main

import "fmt"

//Declaring a function type
type function_type func(i,j int)(int)

func add(i,j int)int{
	return i + j
}

func subtract(i,j int)int{
	return i - j
}


func main(){

	var function_type_instance function_type
	function_type_instance = add
	fmt.Println("Add:",function_type_instance(1,1))

	function_type_instance = subtract
	fmt.Println("Sub:",function_type_instance(1,1))

	fmt.Println("Addding using function factory:",function_factory("add")(1,1))
	fmt.Println("Subtracting using function factory:",function_factory("sub")(23,123))

}

func function_factory(str string) (func(i,j int)(int)){
	switch str {
		case "add":return add
		case "sub":return subtract
		default:return nil
	}

}
