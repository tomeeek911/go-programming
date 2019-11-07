package main

import "fmt"

func main() {
	x := 1
	y := 16
	z := 10
	t := z << 1
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)
	fmt.Printf("%b\n", t)

}

/// << 1 sprawia, ze binarnie przesuwamy dodajemy na koncu jedno zero (przesuwamy binarnie w lewo) np.
