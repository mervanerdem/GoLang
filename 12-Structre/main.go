package main

import "fmt"

type employee struct { // veri tip haline dönüştürerek defalarca struct yaratmaktan kaçınılmaktadır.
	name      string //Aynı zamanda Underlying type diye de geçer
	age       int    //main fonksiyonun içinde de tanımlanabilir
	isMarried bool
	lessons   []string
}

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

	var e1 employee
	e1.name = "Mervan"
	e1.age = 24
	e1.isMarried = false
	e1.lessons = []string{ // Farklı şekilde yazım
		"Math",
		"History",
	}

	var e2 employee
	e2.name = "Ahmet"
	e2.age = 12
	e2.isMarried = false
	e2.lessons = []string{"Math", "English"}

	e3 := employee{ // Bu şekilde de bir kullanım yapılabiilir.
		name:      "Ibrahim",
		age:       30,
		isMarried: true,
		lessons: []string{ // Farklı şekilde yazım
			"English",
			"History",
		},
	}

	fmt.Printf("%#v\n%v\n", e1, e1)
	fmt.Printf("\n%#v\n%v\n", e2, e2)
	fmt.Printf("\n%#v\n%v\n", e3, e3)

}
