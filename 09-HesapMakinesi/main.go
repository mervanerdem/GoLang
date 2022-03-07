package main

import (
	"fmt"
	"gtin"
)

func main() {
	fmt.Print("İşlem yapacağınız ilk sayı:")
	bir, _ := gtin.GetFloat()

	fmt.Print("İşleminiz(+, -, *, /):")
	islem, _ := gtin.GetString()
	for {
		if islem == "+" || islem == "-" || islem == "*" || islem == "/" {
			break
		} else {
			fmt.Println("Yanlış işlem değeri girdiniz. Lütfen tekrar deneyiniz.")
			fmt.Print("İşleminiz(+, -, *, /):")
			islem, _ = gtin.GetString()
		}
	}

	fmt.Print("İşlem yapacağınız ikinci sayı:")
	iki, _ := gtin.GetFloat()

	var sonuc float64 = 0

	switch islem {
	case "+":
		sonuc = bir + iki
		fmt.Println("Sonucunuz:", sonuc)
	case "-":
		sonuc = bir - iki
		fmt.Println("Sonucunuz:", sonuc)
	case "*":
		sonuc = bir * iki
		fmt.Println("Sonucunuz:", sonuc)
	case "/":
		sonuc = bir / iki
		fmt.Println("Sonucunuz:", sonuc)
	default:
		fmt.Println("Something Wrong.")
	}

}
