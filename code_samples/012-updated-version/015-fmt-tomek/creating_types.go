package main

import "fmt"

var a int = 11

type stworzony float64

var b stworzony = 12.1

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

}
