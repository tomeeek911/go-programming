package main

import "fmt"

func main() {

	switch {
	case (2 == 4):
		fmt.Println("chyba nie")
	case (2 == 2):
		fmt.Println("powinno się pokazać")
		fallthrough
	case (1 != 0):
		fmt.Println("powinno pokazac sie po dodaniu fallthrough")
	}
}
