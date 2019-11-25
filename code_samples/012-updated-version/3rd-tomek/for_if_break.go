package main

import "fmt"

func main() {
	x := 5
	for {
		if x > 10 {
			break
		}
		fmt.Println(x)
		x++
	}
}
