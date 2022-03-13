package main

import "fmt"

func main() {
	/* x := 5                      //Pass by value
	slc := []int{0, 1, 2}       //Pass by referances
	myArr := [4]int{0, 1, 2, 3} //Pass by value
	fmt.Println("main1:", x)
	fmt.Println("main1:", slc)
	fmt.Println("main1:", myArr)
	double(x)
	doubleSlc(slc)
	doubleArr(myArr)
	fmt.Println("main2:", x)     //x'in değişmemesinin sebebi değeri fonksiyona kopyalanır.
	fmt.Println("main2:", slc)   //slc'in referans adresi kopyalanır. Slice'ın pointer özelliğinden dolayı.
	fmt.Println("main2:", myArr) // Array'in değeri kopyalanır. O yüzden değer değişmişyor. */

	y := 8
	fmt.Println("main1:", y)
	doublePoi(&y)            //Adresi fonksiyona gönderildi.
	fmt.Println("main2:", y) //adresindeki değeri fonksiyonda 2ile çarptığımız için ana fonksitondaki da değişir.

}

/* func double(x int) {
	x = x * 2
	fmt.Println("Fonksiyon:", x)
}

func doubleSlc(slc []int) {
	for i := 0; i < len(slc); i++ {
		slc[i] *= 2
	}
	fmt.Println("Fonksiyon:", slc)
}

func doubleArr(myArr [4]int) {
	for i := 0; i < len(myArr); i++ {
		myArr[i] *= 2
	}
	fmt.Println("Fonksiyon:", myArr)
} */

func doublePoi(y *int) { //Pointer method
	*y *= 2 //adresindeki değeri 2 ile çarpıyoruz.
	fmt.Println("Fonksiyon:", *y)
}
