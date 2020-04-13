package main

import ("fmt"
	"time"
)

func main(){
	today := time.Now().Weekday()
	switch time.Saturday{
	case today + 0:
		fmt.Println("toda")
	case today + 1:
		fmt.Println("tomorrow")
	default:
		fmt.Println("too far")
	}
	fmt.Println(time.Now())
}