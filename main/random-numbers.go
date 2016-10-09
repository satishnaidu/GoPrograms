package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	f := fmt.Print

	f(rand.Intn(100),",")
	f(rand.Intn(100))
	fmt.Println()

	f(rand.Float64())
	f((rand.Float64()*5)+5,",")
	f((rand.Float64()*5)+5)
	fmt.Println()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	f(r1.Intn(100),",")
	f(r1.Intn(100))
	fmt.Println()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)

	f(r2.Intn(100),",")
	f(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	f(r3.Intn(100), ",")
	f(r3.Intn(100))


}