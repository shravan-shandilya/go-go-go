package main

import "fmt"
import "math/rand"

func main(){
	fmt.Println("replacing a whole bunch of 'if else' ladder with a single nice switch")
	i := rand.Intn(1000)
	switch {
		default:
			fmt.Println("location of 'default' doesnot matter,",i,"is not divisible by 3 or 5")
		case (i % 5 == 0),(i % 3 == 0):
			fmt.Println(i," divisible by both 5 and 3.thanks for supporting musltiple expressions inside of switch")
		case i % 5 == 0:
			fmt.Println(i," divisible by 5,evaluation of expression for switch is awesome")
		case i % 3 == 0:
			fmt.Println(i," divisible by 3,evaluation of expression for switch is awesome")
	}


}
