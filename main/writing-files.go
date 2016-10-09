package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
)

func check (e error){
	if e != nil {
		panic(e)
	}
}

func main(){

	path := "C:/Users/satandey/Desktop/gowrite.txt"
	d1 := []byte("hello\ngo\n")

	err := ioutil.WriteFile(path,d1,0644)
	check(err)

	f, err := os.Create(path)
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
