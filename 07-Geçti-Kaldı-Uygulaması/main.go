//Kullanıcı tarafından girilen nota bağlı olarak dersten geçmesi kontrol edilecek.
//Yapımında Fonksiyon kullanılmalıdır.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Notunuzu Giriniz:")
	not, _ := getInput() //Fonksiyondan not girilmesi istendi.
	if not < 55 {        //Not karşılaştırması
		fmt.Println("Kaldınız.")
	} else {
		fmt.Println("Geçtiniz.")
	}

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
