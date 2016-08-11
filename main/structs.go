package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"satish",29})
	fmt.Println(person{name:"kumar",age:22})
	fmt.Println(person{name:"Fred"})
	fmt.Println(person{age:20})
	//fmt.Println(person(20)) //cannot use 20 (type int) as type string in field value
	s:= person{name:"Sean", age:50}
	fmt.Println(s.name)
	s.age = 55  // Structs are mutable
	fmt.Println(s.age)

}