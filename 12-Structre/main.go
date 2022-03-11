package main

import "fmt"

func main() {

	var employee struct {
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

	fmt.Println(employee)

}
