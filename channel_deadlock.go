package main

import "fmt"

func main(){
  ch := make(chan string)
  
  //Channel sending and recieving are blocking calls.
  //Goruntime recognises that main goroutine(which is the only one) 
  //is sleeping and considers it as a deadlock.
  ch <- "Hello, world"
  
  fmt.Println(<-ch)
}
