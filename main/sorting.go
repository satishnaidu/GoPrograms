package main

import (
	"fmt"
	"sort"
)

func main(){

	strs := []string{"c","a","b"}

	sort.Strings(strs)
	fmt.Println("Strings:",strs)

	ints := []int{7,4, 2}
	sort.Ints(ints)
	fmt.Println("Ints:",ints)

	fmt.Println("Is strings sorted",sort.StringsAreSorted(strs))

}
