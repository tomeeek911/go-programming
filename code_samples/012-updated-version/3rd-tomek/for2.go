package main

import "fmt"

func main() {
	{
		for x := 1; x <= 10; x++ {
			fmt.Printf("%d\t", x)
		}
		for y := 3; y < 7; y++ {
			fmt.Printf("%d\n", y)
		}
	}
}
