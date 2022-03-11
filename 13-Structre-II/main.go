package main

import "fmt"

type employee struct { //Is a Relation Kullanılır:
	name      string //Ahmet ile mehmet işçidir. fakat ahmetle mehmet aynı değildir.
	age       int    //Ya da Aslan ile Kaplan hayvandır hayvan özelliği taşır ama aynı değiller
	isMarried bool
}

type manager struct { //Has a relation kullanılır:
	employee          //Manager bir employe değildir sadece bazı özelliklerini taşır.
	profession string //mantar bir bitki değildir ama bitkilerin bazı özelliklerini taşır.
}

func main() {
	e1 := employee{
		name:      "Mehmet",
		age:       46,
		isMarried: true,
	}

	/* fmt.Println(e1)
	e2 := e1
	e2.name = "Selami" //e1 değerini değiştirmez.
	fmt.Println(e1)
	fmt.Println(e2)
	*/

	m1 := manager{
		employee: employee{
			name:      "Mervan",
			age:       24,
			isMarried: false,
		},
		profession: "Engineer",
	}

	m2 := manager{} // Farklı bir tanımlama şekli
	m2.name = "Salih"
	m2.age = 38
	m2.isMarried = true
	m2.profession = "Buisness Manager"

	fmt.Println(e1)
	fmt.Println(m1)
	fmt.Println(m2)

	CEO := struct { //Anonim Struct yapısı olarak geçer.
		name  string //Birden çok kez kullanılmayacaksa fonksiyon içerisinde kullanılır. Global değişken yapmaya gerek yoktur.
		money int
	}{name: "Cevdet", money: 120000}

	fmt.Println(CEO)

}
