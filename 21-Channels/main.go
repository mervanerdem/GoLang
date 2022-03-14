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
/*
//Burada hata alacağız
package main

import "fmt"

func main() {
	chan1 := make(chan string)

	chan1 <- "Merhaba" //değer kanala gönderildi. başka bir go routine gönderilmek adına

	fmt.Println(<-chan1) // fakat başka bir goroutineden geri dönüş olmadığından foksiyon hataya düşmekte.
} */

package main

import (
	"fmt"
	"math"
)

func main() {
	c1 := circle{5}

	chan1 := make(chan float64)

	go area(c1, chan1)            //fonksiyona c1 ve kanal gönderiliyor.
	fmt.Println("Alan:", <-chan1) //kanal ile alınan değer yazdırılıyor.

}

type circle struct {
	r float64
}

func area(c circle, chan1 chan float64) {
	result := math.Pi * c.r * c.r
	chan1 <- result //return yerine bu şekilde <- ok operatörü kullanılmakta.
}
