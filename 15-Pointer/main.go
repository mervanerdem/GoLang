package main

import "fmt"

func main() {
	name := "Mervan"

	fmt.Println(name)
	fmt.Println(&name) //& --->pointer adres gösterici(Address Operator)

	x := 15
	fmt.Println(x)
	fmt.Println(&x)

}
