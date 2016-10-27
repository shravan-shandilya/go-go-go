package main

import "fmt"
import "math/rand"

func main(){

	fmt.Println("standard conditional")
	i := rand.Intn(100)
	if i%5 == 0 {
		fmt.Println("Inside 'if' and",i,"is divisible by 5")
	}else if i%3 == 0 {
		fmt.Println("Inside 'else if' and",i,"is divisible by 3 not by 5") 
	}else{
		fmt.Println("Inside 'else' and is not divisible by both 5 and 3")
	}

	fmt.Println("advanced conditional 'if' which has initialisation as a part of it")
	if i := rand.Intn(100);i%5 == 0 {
		fmt.Println("Inside 'if' and",i,"is divisible by 5")
	}else if i%3 == 0 {
		fmt.Println("Inside 'else if' and",i,"is divisible by 3 not by 5") 
	}else{
		fmt.Println("Inside 'else' and is not divisible by both 5 and 3")
	}


}
