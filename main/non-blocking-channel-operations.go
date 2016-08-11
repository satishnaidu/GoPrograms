package main

import "fmt"

func main(){

	message := make(chan string)
	signals := make(chan bool)

	select {
	case msg1 := <- message:
		fmt.Println("recevied message "+ msg1)
	default:
		fmt.Println("no message recevied")
	}

	msg := "hai"
	select {
	case message <- msg:
		fmt.Println("sent message ",msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <- message:
		fmt.Println("recevied message "+msg)
	case sig := <- signals:
		fmt.Println("recevied signal ",sig)
	default:
		fmt.Println("no activity")
	}

}
