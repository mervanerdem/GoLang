package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Print("Tahmin edeceğiniz min değeri giriniz:")
	min, err := getInput()
	checkError(err)
	fmt.Print("Tahmin edeceğiniz max değeri giriniz:")
	max, err := getInput()
	checkError(err)
	num := randNum(min, max)
	fmt.Print("Tahmin ettiğiniz değeri giriniz: ")
	guess, err := getInput()
	checkError(err)
	for i := 0; i < 10; i++ {
		if i < 9 {
			if num > guess {
				fmt.Print("Tahmin ettiğiz sayı daha küçük değerden. Tekrar deneyiniz: ")
				guess, err = getInput()
				checkError(err)
			} else if num < guess {
				fmt.Print("Tahmin ettiğiz sayı daha büyük değerden. Tekrar deneyiniz: ")
				guess, err = getInput()
				checkError(err)
			} else {
				fmt.Println("Doğru Tahmin ettiniz. ")
				break
			}
		} else {
			fmt.Println("Bilemediniz.")
		}
	}

}

func randNum(min, max int) int { //Rastgele numara bulma fonksiyonu
	rand.Seed(time.Now().Unix())      //zamanda herhangi bir anda bir zamanı durdurur ve buna bağlı bir değer üretir.
	return (rand.Intn(max-min) + min) //üretilen değeri istenen aralığa göre modunu alır ve return eder.
}

func getInput() (int, error) {
	reader := bufio.NewReader(os.Stdin) //Klavyeden girileni okumak için reader değişkeni atandı.

	input, err := reader.ReadString('\n') //Atanan readerdan stringi okuması istendi
	checkError(err)                       //Hata varsa error verecek fonksiyona yönlendirildi
	input = strings.TrimSpace(input)      //İnput dönüşütürülebilir stringe çevirildi. UniCode dönüşümü. yandaki boşluklar atıldı
	num, err := strconv.Atoi(input)       //String -> İnt dönüşümü yapıldı.
	checkError(err)

	return num, nil //girilen değer cevap olarak gönderildi.
}

func checkError(err error) { //program dönüşünden olacak hata kontrol fonksiyonu
	if err != nil { // error boşluk değerinden farklı mı
		panic(err) //error oluşma durumunda yazdır.
	}
}
