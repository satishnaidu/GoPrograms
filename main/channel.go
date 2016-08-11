package main

import "fmt"

func main(){

	messages := make(chan string)
	consumer := make(chan string)


	go func(){messages <- "ping"}()
	go func(){
		m := <-messages
		consumer <- m
	}()

	msg := <- consumer
	fmt.Println(msg)

}
