//1-Underlying type 'int' olacak şekilde kendi veri tipinizi oluşturunuz
//veri tipi değerini '10' yazdırınız.

package main

import "fmt"

type sayi int

func main() {
	var x sayi = 10
	fmt.Println(x)
}
