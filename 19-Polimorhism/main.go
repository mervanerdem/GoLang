package main

import "fmt"

func main() {
	u := ucgen{3, 4}
	k := kare{8}
	d := dikdortgen{8, 5}

	printArea(u, k, d) //Bu teknikle tek seferde istediğimiz tüm değişkenleri gönderip ayrı ayrı sonuç alabiliriz.

}

type shape interface {
	area() float64
}

func printArea(s ...shape) { //Polimorhism kısmı burda tanımlanıyor.
	for _, shape := range s {
		fmt.Println(shape.area())
	}
}

type ucgen struct {
	a float64
	h float64
}

func (u ucgen) area() float64 {
	fmt.Print("Ucgen Alanı:")
	return u.a * u.h / 2
}

type kare struct {
	a float64
}

func (k kare) area() float64 {
	fmt.Print("Kare Alanı:")
	return k.a * k.a
}

type dikdortgen struct {
	a, b float64
}

func (d dikdortgen) area() float64 {
	fmt.Print("Dikdörtgen Alanı:")
	return d.a * d.b
}
