package main

import "fmt"

func main() {

	for x := 1; x < 5; x++ {
		for y := 1; y <= 10; y++ {
			fmt.Printf("wyniki dla x: %d\t", x)
			fmt.Printf("wyniki dla y: %d\n", y)
		}
	}
}
