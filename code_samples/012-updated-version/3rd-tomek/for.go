package main

import "fmt"

func main() {
	x := 3
	for x <= 11 {
		fmt.Println(x)
		x++
	}
	{
		fmt.Println("Done.")
	}
}
