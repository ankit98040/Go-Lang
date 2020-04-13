package main

import "fmt"

func add(x int, y int) (int, int) {
	out1 := x + y
	var out2 = x - y
	return out1, out2
}

func main() {
	var num1 = 6
	var num2 = 9
	result1, result2 := add(num1, num2)
	fmt.Println(result1, result2)
}
