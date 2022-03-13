//1-Underlying type 'int' olacak şekilde kendi veri tipinizi oluşturunuz
//veri tipi değerini '10' yazdırınız.

/* package main

import "fmt"

type sayi int

func main() {
	var x sayi = 10
	fmt.Println(x)
} */

/* //2-Başlangıç değeri 10 olan bir x değeri atayın.
//Bu x'in adresini y'ye tanımlayıp x değerini 20 diye güncelleyin.

package main

import "fmt"

func main() {
	x := 10
	y := &x
	*y = 20
	fmt.Println(x)
}
*/
//3-Underlying Type struct olan Rectangle type oluşturunuz.
//display,area,circumference metodlarını yazınız.

package main

import "fmt"

type rectangle struct {
	a, b int
}

func (r rectangle) display() {
	fmt.Println("Kenar uzunlukları:", r.a, "ve", r.b, "olan dikdörtgenin")
}

func (r rectangle) area() int {
	return r.a * r.b

}

func (r rectangle) circumference() int {
	return 2 * (r.a + r.b)

}

func main() {
	r1 := rectangle{5, 8}
	r1.display()
	fmt.Println("Alanı:", r1.area())
	fmt.Println("Çevresi:", r1.circumference())
}

//4- 4 adet user'ı struct yapısıyla farklı şekilde tanımlayınız.
//name, age, int -> For döngüsüyle kullanıcıları gösteriniz.
