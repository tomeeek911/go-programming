package main

import "fmt"

const (
	a = iota
	b
	c = iota
	d
	e = 42
	f = iota
	g
	h = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

}
