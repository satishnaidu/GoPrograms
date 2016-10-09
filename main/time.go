package main

import (
	"fmt"
	"time"
)

func main(){
	f := fmt.Println

	now := time.Now()
	f(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	f(then)
	f(then.Year())
	f(then.Month())
	f(then.Day())
	f(then.Hour())
	f(then.Minute())
	f(then.Second())
	f(then.Nanosecond())
	f(then.Unix())
	f(then.UnixNano())
	f(then.Location())

	f(then.Weekday())

	f(then.Before(now))
	f(then.After(now))
	f(then.Equal(now))

	diff := now.Sub(then)
	f(diff)

	f(diff.Hours())
	f(diff.Minutes())
	f(diff.Seconds())
	f(diff.Nanoseconds())

	f(then.Add(diff))
	f(then.Add(-diff))
}
