package main

import "fmt"

var y float64 = 55.001
var z string = "To jest string w nawiasie"
var r string = `Ja mówiłem "To jest string w nawiasie"`
var x int = 45

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(r)
	fmt.Printf("%T\n", r)

	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
