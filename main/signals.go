package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main(){

	signs := make(chan os.Signal,1)
	done := make(chan bool,1)
	signal.Notify(signs, syscall.SIGINT,syscall.SIGTERM)

	go func(){
		sig := <-signs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<- done
	fmt.Println("exiting")
}
