package main

import "fmt"

func main(){
	s := make([]string,3)
	fmt.Println("emp",s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set",s)
	fmt.Println("get",s[2])
	fmt.Println("length",len(s))

	s = append(s,"d")
	s = append(s,"e","f")
	fmt.Println("append",s)

	c := make([]string,len(s))
	copy(c,s)
	fmt.Println("copy",c)

	l := s[2:5]
	fmt.Println("sl1:", l)
	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3",l)

	t := []string{"g","h","i"}
	fmt.Println("dc1",t)

	twod := make([][]int,3)
	for i:=0; i<3;i++ {
		innerLen :=i+1
		twod[i] = make([]int,innerLen)
		for j :=0; j<innerLen;j++ {
			twod[i][j] = i + j
		}
	}

	fmt.Println("2d",twod)

}
