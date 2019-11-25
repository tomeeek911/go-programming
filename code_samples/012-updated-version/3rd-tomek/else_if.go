package main

import "fmt"

func main() {
	x := 42
	if x != 42 {
		fmt.Println("wartość jest rózna od 42")
	} else if x == 41 {
		fmt.Println("wartosc jest rowniez nierowna 42, a 41")
	} else {
		fmt.Println("rowne 42")
	}
}
