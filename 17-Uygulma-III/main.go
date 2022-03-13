//1-Underlying type 'int' olacak şekilde kendi veri tipinizi oluşturunuz
//veri tipi değerini '10' yazdırınız.

/* package main

import "fmt"

type sayi int

func main() {
	var x sayi = 10
	fmt.Println(x)
} */

//Başlangıç değeri 10 olan bir x değeri atayın.
//Bu x'in adresini y'ye tanımlayıp x değerini 20 diye güncelleyin.

package main

import "fmt"

func main() {
	x := 10
	y := &x
	*y = 20
	fmt.Println(x)
}
