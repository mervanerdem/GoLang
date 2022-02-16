package main

import "fmt"

func topla(a int, b int) int {
	return a + b //a ve toplamını gönderir.
}

func yazdir() {
	fmt.Println("İyi Günler")
}

func main() {
	fmt.Println(topla(2, 5))
	yazdir()

}
