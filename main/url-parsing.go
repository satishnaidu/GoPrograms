package main

import (
	"fmt"
	"net"
	"net/url"
)

func main(){

	f := fmt.Println
	s := "postgres://user:pass@host.com:5542/path?k=v#f"

	u, err := url.Parse(s)
	if(err != nil){
		panic(err)
	}

	f(u.Scheme)
	f(u.User)
	f(u.User.Username())
	p, _ := u.User.Password()
	f(p)

	f(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	f(host)
	f(port)

	f(u.Path)
	f(u.Fragment)
	f(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	f(m)
	f(m["k"][0])





}
