package main

import "fmt"
import "time"

func main(){

	i := 2
	fmt.Println("write",i,"as")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday,time.Sunday:
		fmt.Println("its a weekend")
	default:
		fmt.Println("Its a weekday")
	}

	t:= time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
