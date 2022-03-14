package main

import (
	"fmt"
	"time"
)

func main() {
	//Go Routine aynı anda birden çok işlemin yapılmasını sağlayan yapıdır.

	go printX()
	go printY()
	time.Sleep(time.Second)    //
	fmt.Println("Main Bitişi") // Çalışmasını bitirmeden main GoRoutini bittiğinden yalnızca çıktı görülmekte.
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
