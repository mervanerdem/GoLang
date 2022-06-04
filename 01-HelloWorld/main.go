//Hello World
package main /*Package, kod sayfalarımız arasında iletişimde bulunabilmemizi sağlar.
Bu sayede içerisinde package değeri aynı olan kod sayfaları birbirleriyle iletişim halinde
olma yeteneği kazanır. package uygulama olan sayfalar birbiriyle iletişim kurabilir.*/
import "fmt" //Ana işlemlerin yapıldığı kütüphane

func main() { //Programın çalışacağı ana fonksiyon.
	fmt.Println("Hello World") //fmt kütüphanesinin içindeki print(ekrana yazdırma) ile ekrana yazdırma
}
