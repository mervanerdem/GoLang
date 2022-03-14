package main

import "fmt"

func main() {
	/* 	var myChannel chan string
	   	myChannel = make(chan string) //Uzun yöntem channel tanımlama
	*/
	myChannel := make(chan string) //Kısa yol channel tanımlama
	go merhaba(myChannel)

	fmt.Println(<-myChannel)
}
func merhaba(channel chan string) {
	channel <- "Merhaba"
}
