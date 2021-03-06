//Yapılmak istenen uygulamanın yorum satırını kaldırınca deneyebileceksiniz.
package main

import "fmt"

func main() {
	/* //1.Örnek) 1 ile 10 arasındaki sayıları tek ve çift olarak yazdırın
	for i := 1; i <= 10; i++ {
		switch i % 2 {
		case 0:
			fmt.Println(i, "sayısı çift sayıdır.")
		case 1:
			fmt.Println(i, "sayısı tek sayıdır.")

		}

	} */

	/* 	//2.Örnek) For yapısını kullanarak Go dilinde olmayan While Döngüsüne örnek veriniz.
	   	i := 10
	   	for true {
	   		i--
	   		fmt.Println(i)
	   		if i <= 0 {
	   			break
	   		}
	   	}
	*/
	/* 	//3.Örnek)Switch FallThrough ifadesini açıklayınız.
	   	//fallThrough ifadesi sorgu bittiğinde break yerine diğer ifadeleri de sorgulamasını istemekte.

	   	switch i := 3; {
	   	case i < 5:
	   		fmt.Println(i, "5'ten küçük")
	   		fallthrough
	   	case i < 10:
	   		fmt.Println(i, "10'ten küçük")
	   		fallthrough
	   	case i < 15:
	   		fmt.Println(i, "15'ten küçük")
	   		fallthrough
	   	case i < 20:
	   		fmt.Println(i, "20'ten küçük")
	   	} */

	//4.Örnek) 1 ile 50 arasındaki asal sayıları gösteren program
	var x, y int
	for x = 2; x <= 50; x++ {
		for y = 2; float64(y) < float64(x)/float64(y); y++ {
			if x%y == 0 {
				break
			}
		}
		if y > x/y {
			fmt.Println(x, "bir asal sayıdır.")
		}
	}
}
