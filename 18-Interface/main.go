package main

import (
	"fmt"
	"math"
)

func main() {
	r1 := rectangle{5, 8}
	/* 	fmt.Println("Alanı:", r1.area())
	   	fmt.Println("Çevresi:", r1.circumference())
	fmt.Println("-----------")*/
	interfaceFunc(r1)

	fmt.Println("-----------")
	c1 := circle{5.0}
	interfaceFunc(c1)

}

type shape interface {
	area() float64
	circumference() float64
}

func interfaceFunc(i shape) {
	fmt.Println("i:", i)
	fmt.Println("Alan:", i.area())
	fmt.Println("Çevre:", i.circumference())
	fmt.Printf("%T\n", i)
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c circle) circumference() float64 {
	return 2 * math.Pi * c.r
}

type rectangle struct {
	a, b float64
}

func (r rectangle) area() float64 {
	return r.a * r.b

}

func (r rectangle) circumference() float64 {
	return 2 * (r.a + r.b)

}
