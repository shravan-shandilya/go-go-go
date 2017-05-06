package main

import (
	"fmt"
)

func main() {
	var temp []int

	for i:=0;i<25;i++ {
		temp = append(temp,i)
		fmt.Println(cap(temp),temp)
		if i == 25 {
			break
		}
	}
}
