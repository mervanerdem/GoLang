package main

import "fmt"

func main() {
	/* 	name := "Mervan"

	   	fmt.Println(name)
	   	fmt.Println(&name) //& --->pointer adres gösterici(Address Operator) */

	//x := 15
	/* 	fmt.Printf("%T  %v\n", x, x)   //%T veri tipini, \n alt satıra geçişi ve %v değerini ifade etmektedir.
	   	fmt.Printf("%T  %v\n", &x, &x) //pointer adreslerini görüntüleme

	   	y := &x                      //y pointer olarak tanımlanıp x'in adresi atandı.
	   	fmt.Printf("%T  %v\n", y, y) //y *int değişkenine sahip */
	/*
		fmt.Println(x)        // x değeri
		fmt.Println(&x)       //x'in adresi
		fmt.Println(*(&x))    //x'in adresinin gösterdiği değer-> yani x
		fmt.Println(&(*(&x))) // //x'in adresinin gösterdiği değerin adresi yani-> x'in adresi
		// * adresteki değeri gösterir, & değerin adresini gösterir. Yani birbirine dönüşüm yapılabilir. */

	x1 := 12
	x2 := &x1

	fmt.Println(x1, x2)
	fmt.Println(x1, *x2) // adres gösteren x2 adresin gösterdiği değeri gösterdi o da x'e eşit

	*x2 = 8 // adresteki değere 8 sayısını atadık

	fmt.Println(x1, *x2) // bu şekilde x'in de değeri de değişmiş oldu.

	x3 := &x1
	*x3 = 92

	fmt.Println(x1, *x2, *x3)

}
