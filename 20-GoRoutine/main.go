package main

import (
	"fmt"
)

func main() {

	go printX()
	go printY()
	fmt.Println("Main Bitişi")
}

func printX() {
	for i := 0; i < 10; i++ {
		fmt.Print("X")
	}
	fmt.Println()
	fmt.Println("---------")
}

func printY() {
	for i := 0; i < 10; i++ {
		fmt.Print("Y")
	}
	fmt.Println()
	fmt.Println("---------")
}
