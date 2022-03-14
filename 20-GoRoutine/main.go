package main

import (
	"fmt"
)

func main() {

	printX()
	printY()

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
