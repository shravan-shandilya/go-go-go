package main

import "fmt"

func main(){
  //Unbuffered channel, it means that channel cannot store any messages.
  //Every message sent has to be immediately transfered to reciever
  ch := make(chan string)
  
  //Channel sending and recieving are blocking calls.
  //Goruntime recognises that main goroutine(which is the only one) 
  //is sleeping and considers it as a deadlock.
  ch <- "Hello, world"
  
  fmt.Println(<-ch)
}
