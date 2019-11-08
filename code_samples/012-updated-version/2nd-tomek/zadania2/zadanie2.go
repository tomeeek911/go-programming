package main

import "fmt"

var (
	a = 1
	b = 2
	c = 3
)

func main() {
	fmt.Println(a == b)
	fmt.Println(a <= b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)
}
