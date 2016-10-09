package main

import (
	"fmt"
	"flag"
)

func main(){

	wordPtr := flag.String("word","foo","a string")

	numPtr := flag.Int("numb",7,"a number")
	boolPtr := flag.Bool("fork",false,"a boolean")

	var svar string
	flag.StringVar(&svar,"svar","bar","a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())


}