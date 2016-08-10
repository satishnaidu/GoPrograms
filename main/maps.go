package main

import "fmt"

func main(){
	m:= make(map[string] int)
	fmt.Println(m)
	m["k1"] =7
	m["k2"] = 10
	fmt.Println(m)

	v1 :=m["k1"]
	fmt.Println("v1",v1)

	fmt.Println("length",len(m))

	delete(m,"k2")
	fmt.Println("map",m)

	_, prs := m["k2"]
	fmt.Println("prs",prs)

	n := map[string]int{"foo":1,"bar":2}
	fmt.Println("map",n)

}
