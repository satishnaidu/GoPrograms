package main

import (
	"fmt"
	"sync/atomic"
	"runtime"
	"time"
)

func main(){

	var ops uint64 = 0

	for i := 1; i<=50; i++ {
		go func(){
			for {
				atomic.AddUint64(&ops,1)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println(opsFinal)
}
