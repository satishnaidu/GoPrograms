package main

import "fmt"
import "time"

func main(){
	c1 := make(chan string,1)
	go func(){
	        time.Sleep(time.Second*2)
		c1 <- "result 1"
	}()

	select {
	case msg := <- c1:
		fmt.Println(msg)
	case <-time.After(time.Second*1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string)
	go func(){
		time.Sleep(time.Second*2)
		c2 <- "result 2"
	}()

	select {
	case msg := <- c2:
		fmt.Println(msg)
	case <-time.After(time.Second*4):
		fmt.Println("timeout 2")
	}
}
