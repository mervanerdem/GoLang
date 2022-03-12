package main

import "fmt"

type mile float64
type kilometer float64

func main() {

	var m1 mile = 5.7
	fmt.Printf("%T\n%v\n", m1, m1)

	/* 	m2 := float64(8.4)
	   	fmt.Printf("\n%T\n%v\n", m2, m2)

	   	m3 := m1 + mile(m2) //veri tipleri farklı olduğundan m2 yi dönüştürmemiz gerekmektedir.
	   	fmt.Printf("\n%T\n%v\n", m3, m3)
	*/

	m2 := mile(1.2)
	fmt.Println("m1+m2:", m1+m2) // Aynı veri tipinde olduğundan toplanmasında bir mahsur bulunmamakta.

	k1 := kilometer(10.1)

	fmt.Printf("%T\n%v\n", k1, k1)

}
