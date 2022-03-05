//Yapılmak istenen uygulamanın yorum satırını kaldırınca deneyebileceksiniz.
package main

import "fmt"

func main() {
	//1.Örnek) 1 ile 10 arasındaki sayıları tek ve çift olarak yazdırın
	for i := 1; i <= 10; i++ {
		switch i % 2 {
		case 0:
			fmt.Println(i, "sayısı çift sayıdır.")
		case 1:
			fmt.Println(i, "sayısı tek sayıdır.")

		}

	}
}
