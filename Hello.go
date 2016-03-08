package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5000; i++ {
		runtime.Gosched()
		fmt.Println(s)
		fmt.Print(i)
	}
}

func main() {
	go say("world")
	say("hello")
}
