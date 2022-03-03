package main

import "fmt"

func main() {
	x := 5   //int x tanımlaması
	y := 5.5 //float y tanımlaması

	fmt.Printf("%v %T\n", x, x) //veri tipi ve değeri sorgulama
	fmt.Printf("%v %T\n", y, y) //  \n anlamı bir alt satıra geç demektir.

	z := float64(x) + y //aynı tip olmadığından aynı tipe dönüştürme
	fmt.Println(z)
	t := x + int(y) // bu şekilde de veri dönüştürme yapılabilir.
	fmt.Println(t)

}

/* SONUÇ
5 int		// x değeri
5.5 float64	// y değeri
10.5		//z değeri
10			//t değeri
*/
