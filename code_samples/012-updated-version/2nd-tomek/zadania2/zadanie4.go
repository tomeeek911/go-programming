package main

import "fmt"

var x int = 44
var y = x << 1

func main() {
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\n", x)
	y := x << 1
	fmt.Printf("%d\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%#x\n", y)
}
