package main

import "fmt"

type rect2 struct{
	width, height int
}

func (r *rect2) area() int {
	return r.width * r.height
}

func (r rect2) perim() int {
	return 2*r.width + 2*r.height
}


func main() {
	r := rect2{width:10, height:5}
	fmt.Println("area ",r.area())
	fmt.Println("perim ",r.perim())

	rp := &r
	fmt.Println("area ",rp.area())
	fmt.Println("perim ",rp.perim())
}