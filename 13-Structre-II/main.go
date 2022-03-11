package main

import "fmt"

type employee struct { // veri tip haline dönüştürerek defalarca struct yaratmaktan kaçınılmaktadır.
	name      string //Aynı zamanda Underlying type diye de geçer
	age       int    //main fonksiyonun içinde de tanımlanabilir
	isMarried bool
}

func main() {
	e1 := employee{
		name:      "Mehmet",
		age:       46,
		isMarried: true,
	}
	fmt.Println(e1)
	e2 := e1
	e2.name = "Selami" //e1 değerini değiştirmez.

	fmt.Println(e1)
	fmt.Println(e2)

}
