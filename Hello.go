package main

import (
	"fmt"
	"github.com/saadjaved120/stringUtil"
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
	say(stringUtil.Reverse("HELLO There"))
}
