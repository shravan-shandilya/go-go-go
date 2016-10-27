package main

import "fmt"

func main(){

	fmt.Println("standard for loop")
	for i := 0;i<10;i++ {
		fmt.Println(i)
	}

	fmt.Println("'while' styled for loop")
	j := 10
	for j > 0{
		fmt.Println(j)
		j--
		if j == 5 {
			j--
			fmt.Println("this is how break works")
			break
		}

	}




/*
	fmt.Println("The mighty infinite loop")
	for {
		fmt.Println("inside of nowhere")
	}
*/
}
