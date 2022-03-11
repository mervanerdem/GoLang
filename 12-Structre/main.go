package main

import "fmt"

func main() {
	//Temel Struct Yapısı
	/* 	var employee struct {
	   		name      string
	   		age       int
	   		isMarried bool
	   	}

	   	fmt.Printf("%#v \n%v\n", employee, employee)
	   	fmt.Println("-----------")

	   	employee.name = "Mervan"
	   	fmt.Println(employee.name)
	   	employee.age = 24
	   	employee.isMarried = false

	   	fmt.Println(employee)   */

	type employee struct { // veri tip haline dönüştürerek defalarca struct yaratmaktan kaçınılmaktadır.
		name      string //Aynı zamanda Underlying type diye de geçer
		age       int
		isMarried bool
	}

	var e1 employee
	e1.name = "Mervan"
	e1.age = 24
	e1.isMarried = false

	var e2 employee
	e2.name = "Ahmet"
	e2.age = 12
	e2.isMarried = false

	fmt.Printf("%#v\n%v\n", e1, e1)
	fmt.Printf("%#v\n%v\n", e2, e2)

}
