package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main(){

	f := fmt.Println
	bolB, _ := json.Marshal(true)
	f(string(bolB))

	intB, _ := json.Marshal(1)
	f(string(intB))
	fltB, _ := json.Marshal(2.13)
	f(string(fltB))
	strB, _ := json.Marshal("gopher")
	f(string(strB))

	slcD := []string{"apple","peach","pear"}
	slcB, _ := json.Marshal(slcD)
	f(string(slcB))

	mapD := map[string]int {"apple":5,"lettuce":7}
	mapB, _ := json.Marshal(mapD)
	f(string(mapB))

	resD := &Response1{
		Page:10,
	        Fruits:[]string{"apple","peach","pear"}}
	resB, _ := json.Marshal(resD)
	f(string(resB))

	res2D := &Response2{
		Page:1,
		Fruits:[]string{"apple","peach","pear"}}
	res2B, _ := json.Marshal(res2D)
	f(string(res2B))

	byt := []byte(`{"num":6.13,"str":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	f(dat)

	num := dat["num"].(float64)
	f(num)
	strs := dat["str"].([]interface{})
	str1 := strs[0].(string)
	f(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res2 := Response2{}
	json.Unmarshal([]byte(str), &res2)
	f(res2)
	f(res2.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple":5,"lettuce":7}
	enc.Encode(d)
}
