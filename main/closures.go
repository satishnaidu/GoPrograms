package main

import "fmt"


func initSeq() func() int {
	i := 0
	return func() int {
		i +=1
		return i
	}
}

func main(){
	nextInt := initSeq();
	fmt.Println(nextInt)
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	nextInt2 := initSeq();
	fmt.Println(nextInt2)
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
}