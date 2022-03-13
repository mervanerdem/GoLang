//Get Input Version 1.0
//Powered By github.com/mervanerdem

package gtin

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetInt() (int, error) {
	reader := bufio.NewReader(os.Stdin) //Klavyeden girileni okumak için reader değişkeni atandı.

	input, err := reader.ReadString('\n') //Atanan readerdan stringi okuması istendi
	if err != nil {                       // error boşluk değerinden farklı mı
		panic(err) //error oluşma durumunda yazdır.
	} //Hata varsa error verecek fonksiyona yönlendirildi
	input = strings.TrimSpace(input) //İnput dönüşütürülebilir stringe çevirildi. UniCode dönüşümü. yandaki boşluklar atıldı
	num, err := strconv.Atoi(input)  //String -> İnt dönüşümü yapıldı.
	if err != nil {                  // error boşluk değerinden farklı mı
		panic(err) //error oluşma durumunda yazdır.
	}

	return num, nil //girilen değer cevap olarak gönderildi.
}

func GetString() (string, error) {
	reader := bufio.NewReader(os.Stdin) //Klavyeden girileni okumak için reader değişkeni atandı.

	input, err := reader.ReadString('\n') //Atanan readerdan stringi okuması istendi
	if err != nil {                       // error boşluk değerinden farklı mı
		panic(err) //error oluşma durumunda yazdır.
	} //Hata varsa error verecek fonksiyona yönlendirildi
	input = strings.TrimSpace(input)
	return input, nil //girilen değer cevap olarak gönderildi.
}

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin) //Klavyeden girileni okumak için reader değişkeni atandı.

	input, err := reader.ReadString('\n') //Atanan readerdan stringi okuması istendi
	if err != nil {                       // error boşluk değerinden farklı mı
		panic(err) //error oluşma durumunda yazdır.
	} //Hata varsa error verecek fonksiyona yönlendirildi
	input = strings.TrimSpace(input) //İnput dönüşütürülebilir stringe çevirildi. UniCode dönüşümü. yandaki boşluklar atıldı
	num, err := strconv.ParseFloat(input, 64)

	//num, err := strconv.Atoi(input) //String -> İnt dönüşümü yapıldı.
	if err != nil { // error boşluk değerinden farklı mı
		panic(err) //error oluşma durumunda yazdır.
	}

	return num, nil //girilen değer cevap olarak gönderildi.
}
