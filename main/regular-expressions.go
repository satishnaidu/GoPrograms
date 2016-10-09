package main

import (
	"fmt"
	"regexp"
	"bytes"
)

func main(){

	f := fmt.Println
	match, _ := regexp.MatchString("p([a-z]+)ch","peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	f(r.MatchString("peach"))
	f(r.FindString("peach punch"))
	f(r.FindStringIndex("peach punch"))
	f(r.FindStringSubmatch("peach punch"))
	f(r.FindStringSubmatchIndex("peach punch"))
	f(r.FindAllString("peach punch pinch", -1))
	f(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	f(r.FindAllString("peach punch pinch", 2))
	f(r.Match([]byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch")
	f(r)
	f(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in,bytes.ToUpper)
	f(string(out))
}
