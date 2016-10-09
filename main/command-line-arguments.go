package main

import (
	"fmt"
	"os"
)

func main(){

	argsWithProg := os.Args
	argsWithOutProg := os.Args[1:]

	if len(os.Args)> 3 {
		args := os.Args[3]
		fmt.Println(args)
	}

	fmt.Println(argsWithProg)
	fmt.Println(argsWithOutProg)

}