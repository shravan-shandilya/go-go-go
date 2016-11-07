package main

import "fmt"

//Alias or typedef in go
type my_int int


var (
	global_string = "this is a global_string"
	integer_32 int32 = 12
	integer_64 int64 = 34
        float_64 float64 = 40.
	my_int_instance my_int = 56
)

const (
	constant_string = "this is a constant string"	
)
func main(){
	local_string := "string in main"
	fmt.Println("Printing from main (local_string):",local_string)
	fmt.Println("Printing from main (global_string):",global_string)
	test()
	fmt.Println("Printing from main (local_string):",local_string)

	//trying to alter a constant. very bad idea!
	constant_string = "i cant do this,If you get the error,remove this line"
}

func test(){
	local_string := "string in the test"
	fmt.Println("Printing from test (local_string):",local_string)
	fmt.Println("Printing from test (global_string):",global_string)
	global_string = "modified global_string from function test"
	fmt.Println("Printing from test (global_string):",global_string)
}
