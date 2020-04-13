package main

import (
	"fmt"
	"time"
)

func main() {
	//count("sheep")   go routines helps for better concurrency
	go count("sheep")
	count("fish")
}

func count(thing string) {
	for i := 0; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
