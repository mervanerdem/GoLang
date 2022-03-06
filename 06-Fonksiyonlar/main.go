//Fonksiyon işlemleri
package main

import "fmt"

func main() {
	hello() //değişken değeri geri döndürmeyen fonksiyon
	a := 5  // değer ataması
	b := 18
	c := sum(a, b) //toplama yapan fonksiyon
	write(c)       //içine yazılan değeri ekrana yazdıran fonksiyon

}

func hello() { //değer döndürmediği için parantezler içerisine herhangi bir şey yazmıyoruz.
	fmt.Println("Hello Function")
}

func sum(x, y int) int { //parantez içerisindeki x ve y int değerinde ana fonsiyondan gelecek değerlerdir.
	//parantez dışındaki int ise fonksiyon sonucun geri dönüş yapacak değerin tipi
	var t int = x + y //toplama işelmini yapan değişken
	return t          // return ile içerideki fonksiyonu geri döndürüyoruz
	//return x + y // şeklinde de bir yazım yapılabilirdi.
}

func write(x int) { //dışardan bir değer alınacak
	fmt.Println("Değer:", x) //alınan değer doğrudan ekrana yazdırılacak
	//herhangi bir değer döndürmediğinden return komutuna ihtiyaç bulunmamakta.
}
