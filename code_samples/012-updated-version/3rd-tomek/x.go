package main

import "fmt"

func main() {
	for x := 1; x < 25; x++ {
		if x%7 == 3 {
			fmt.Println(x)
		}
	}
}

//x%7 == 3 - daj mi kazda co siodma liczbe zaczynajac od 3 i konczac max na 25
