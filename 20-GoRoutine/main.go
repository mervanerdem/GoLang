package main

import (
	"fmt"
	"time"
)

func main() {
	//Go Routine aynı anda birden çok işlemin yapılmasını sağlayan yapıdır.

	go printX()                //Süre bekletmesi dururken Hhangi işlem daha önce başlıyor hangisi önce bitiyor belirsiz.
	go printY()                //Bazı çıktılarda X önce bastırılırken ekrana bazı çıktılarda ise Y önce bastırılıyor. Sistemde kararsızlık bulunmakta.
	time.Sleep(time.Second)    //Ana Fonksiyonu bir saniye bekletiyoruz diğer işlemler yapılması için.
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
