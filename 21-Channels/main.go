/* package main

import "fmt"

func main() {
	// 	var myChannel chan string
	 //  	myChannel = make(chan string) //Uzun yöntem channel tanımlama

	myChannel := make(chan string) //Kısa yol channel tanımlama
	go merhaba(myChannel)          //myChannel değişkeni GoRoutine ile gönderildi.

	fmt.Println(<-myChannel) //burda da döndürdüğümüz değeri okutabiliyoruz.
}
func merhaba(channel chan string) {
	channel <- "Merhaba" //Buradan ok operatörü ile return yerine fonksiyondan geri dönüş sağladık.
} */

//Burada hata alacağız
package main

import "fmt"

func main() {
	chan1 := make(chan string)

	chan1 <- "Merhaba" //değer kanala gönderildi. başka bir go routine gönderilmek adına

	fmt.Println(<-chan1) // fakat başka bir goroutineden geri dönüş olmadığından foksiyon hataya düşmekte.
}
