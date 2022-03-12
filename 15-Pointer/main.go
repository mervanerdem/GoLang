package main

import "fmt"

func main() {
	/* 	name := "Mervan"

	   	fmt.Println(name)
	   	fmt.Println(&name) //& --->pointer adres gösterici(Address Operator) */

	x := 15
	fmt.Printf("%T  %v\n", x, x)   //%T veri tipini, \n alt satıra geçişi ve %v değerini ifade etmektedir.
	fmt.Printf("%T  %v\n", &x, &x) //pointer adreslerini görüntüleme

	y := &x                      //y pointer olarak tanımlanıp x'in adresi atandı.
	fmt.Printf("%T  %v\n", y, y) //y *int değişkenine sahip

}
