package main

import "fmt"

var x byte

func main() {
	x = 12
	y := 12.1111
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
