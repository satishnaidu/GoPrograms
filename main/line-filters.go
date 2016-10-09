package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(){

		ucl := strings.ToUpper(scanner.Text())
		if ucl == "QUIT" {
			os.Exit(0)
		}
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr,"error:",err)
		os.Exit(1)
	}
}