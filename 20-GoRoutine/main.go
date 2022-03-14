package main

import (
	"fmt"
	"sync"
)

func main() {
	//Go Routine aynı anda birden çok işlemin yapılmasını sağlayan yapıdır.
	wg.Add(1) //burdan sonra 1 tane bekleyecek goroutine olduğunu gösterir.
	go printX()
	wg.Wait() //burda bekleyecek wg.Add kısmında göterilen goroutini
	printY()  //bu şekilde sırayla yapılacak tek çekirdekli işlemciler için işlem önceliği verilmiş olur.

}

var wg sync.WaitGroup // Wait Group oluşturuşturuldu. GoRoutine bitişini bekler.

func printX() {
	for i := 0; i < 10; i++ {
		fmt.Print("X")
	}
	fmt.Println()
	fmt.Println("---------")
	wg.Done() //GoRoutine bittiğini devam edebebilir sinyali
}

func printY() {
	for i := 0; i < 10; i++ {
		fmt.Print("Y")
	}
	fmt.Println()
	fmt.Println("---------")
}
