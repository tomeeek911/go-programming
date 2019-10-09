package main

import "fmt"

func main (){
	fmt.Println("Wyświetlanie tekstu razem z funkcjami wewnetrznymi")
	srodek ()
	fmt.Println(", ktore pokazuja czy cos dziala dobrze, czy nie")
}
func srodek (){
	fmt.Println("(wlasnie tutaj)")
}
