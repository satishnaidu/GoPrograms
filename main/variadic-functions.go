package main

import "fmt"

func sum(nums ...int){
	fmt.Println(nums," ")
	total := 0
	for sum := range nums {
		total += sum
	}

	fmt.Println("total sum",total)

}

func main(){
	sum(1,2)
	sum(2,3,4)
	sum(4,5,6,7)
	nums := []int {1,2,3,4,5,6}
	sum(nums...)
}
